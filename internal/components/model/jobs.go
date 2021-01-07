/*
------------------------------------------------------------------------------------------------------------------------
####### model ####### (c) 2020-2021 mls-361 ######################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package model

import "github.com/mls-361/armen-sdk/jw"

func (cm *Model) newJob(job *jw.Job) {
	var wf string

	if job.Workflow != nil {
		wf = *job.Workflow
	}

	cm.components.Logger.Info( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
		"New job",
		"id", job.ID,
		"name", job.Name,
		"namespace", job.Namespace,
		"type", job.Type,
		"origin", job.Origin,
		"priority", job.Priority,
		"exclusivity", job.Exclusivity,
		"workflow", wf,
	)

	cm.publish("new.job", *job) //**************************************************************************************
}

// InsertJob AFAIRE.
func (cm *Model) InsertJob(job *jw.Job) error {
	done, err := cm.components.Backend.InsertJob(job)
	if err != nil {
		var wf string

		if job.Workflow != nil {
			wf = *job.Workflow
		}

		cm.components.Logger.Error( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
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
		cm.components.Logger.Notice( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"A job with the same key already exists",
			"name", job.Name,
			"namespace", job.Namespace,
			"type", job.Type,
			"unique_key", *job.UniqueKey,
		)

		return nil
	}

	cm.newJob(job)

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
