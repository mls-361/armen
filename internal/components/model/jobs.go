/*
------------------------------------------------------------------------------------------------------------------------
####### model ####### (c) 2020-2021 mls-361 ######################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package model

import (
	"time"

	"github.com/mls-361/armen-sdk/jw"
)

const (
	_njTimeout = 5 * time.Minute
)

func (cm *Model) newJob(job *jw.Job) {
	var wf string

	if job.Workflow != nil {
		wf = *job.Workflow
	}

	cm.mcsJobsCreated.Inc()

	cm.components.Logger().Info( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
		"New job",
		"id", job.ID,
		"name", job.Name,
		"namespace", job.Namespace,
		"type", job.Type,
		"origin", job.Origin,
		"priority", job.Priority,
		"workflow", wf,
	)

	cm.publish("new.job", *job) //**************************************************************************************
}

// InsertJob AFAIRE.
func (cm *Model) InsertJob(job *jw.Job) error {
	done, err := cm.components.CBackend.InsertJob(job)
	if err != nil {
		var wf string

		if job.Workflow != nil {
			wf = *job.Workflow
		}

		cm.components.CLogger.Error( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"Impossible to insert a new job",
			"id", job.ID,
			"name", job.Name,
			"namespace", job.Namespace,
			"type", job.Type,
			"origin", job.Origin,
			"priority", job.Priority,
			"workflow", wf,
			"reason", err.Error(),
		)

		return err
	}

	if !done {
		cm.components.CLogger.Notice( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"A job with the same key already exists",
			"name", job.Name,
			"namespace", job.Namespace,
			"type", job.Type,
			"key", *job.Key,
		)

		return nil
	}

	cm.newJob(job)

	return nil
}

// NextJob AFAIRE.
func (cm *Model) NextJob() *jw.Job {
	cm.njMutex.Lock()
	defer cm.njMutex.Unlock()

	if cm.njTimeout.After(time.Now()) {
		return nil
	}

	job, err := cm.components.CBackend.NextJob()
	if err != nil {
		cm.components.CLogger.Warning( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"Cannot retrieve the next job to run",
			"reason", err.Error(),
		)

		cm.njTimeout = time.Now().Add(_njTimeout)

		return nil
	}

	if job != nil {
		cm.publish("job.before", *job) //*******************************************************************************
	}

	return job
}

// UpdateJob AFAIRE.
func (cm *Model) UpdateJob(job *jw.Job) {
	cm.publish("job.after", *job) //************************************************************************************

	switch job.Status {
	case jw.StatusFailed:
		cm.mcsJobsFailed.Inc()
	case jw.StatusSucceeded:
		cm.mcsJobsSucceeded.Inc()
	}

	if err := cm.components.CBackend.UpdateJob(job); err != nil {
		cm.components.CLogger.Error( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"Impossible to update this job",
			"id", job.ID,
			"name", job.Name,
			"namespace", job.Namespace,
			"type", job.Type,
			"reason", err.Error(),
		)

		return
	}

	if job.Workflow == nil || job.Status == jw.StatusPending {
		return
	}

	wf, err := cm.components.CBackend.Workflow(*job.Workflow, true)
	if err != nil {
		cm.components.CLogger.Error( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"Cannot retrieve the workflow associated with this job",
			"id", job.ID,
			"name", job.Name,
			"namespace", job.Namespace,
			"type", job.Type,
			"workflow", *job.Workflow,
			"reason", err.Error(),
		)

		return
	}

	if err := cm.updateWorkflow(job, wf); err != nil {
		cm.components.CLogger.Error( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"Impossible to update this workflow",
			"id", wf.ID,
			"reason", err.Error(),
		)
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
