/*
------------------------------------------------------------------------------------------------------------------------
####### model ####### (c) 2020-2021 mls-361 ######################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package model

import (
	"time"

	"github.com/mls-361/armen-sdk/jw"
	"github.com/mls-361/failure"
	"github.com/mls-361/uuid"
)

func (cm *Model) stepToJob(wf *jw.Workflow, stepName string) (*jw.Job, error) {
	step, ok := wf.AllSteps[stepName]
	if !ok {
		return nil, failure.New(nil).
			Set("step", stepName).
			Msg("this step does not exist") ////////////////////////////////////////////////////////////////////////////
	}

	job := jw.NewJob(
		uuid.New(),
		stepName,
		step.Namespace,
		step.Type,
		wf.Origin,
		wf.Priority,
		nil,
	)

	job.Workflow = &wf.ID
	job.Emails = wf.Emails
	job.Config = step.Config
	job.TimeReference = wf.CreatedAt

	return job, nil
}

func (cm *Model) firstJob(wf *jw.Workflow) (*jw.Job, error) {
	job, err := cm.stepToJob(wf, wf.FirstStep)
	if err != nil {
		return nil, err
	}

	job.Public = wf.Data

	wfFailed := false
	job.WorkflowFailed = &wfFailed

	return job, nil
}

func (cm *Model) newWorkflow(wf *jw.Workflow) {
	cm.components.CLogger.Info( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
		"New workflow [begin]",
		"id", wf.ID,
		"description", wf.Description,
		"origin", wf.Origin,
		"priority", wf.Priority,
	)

	cm.publish("new.workflow", *wf) //**********************************************************************************
}

// InsertWorkflow AFAIRE.
func (cm *Model) InsertWorkflow(wf *jw.Workflow) error {
	job, err := cm.firstJob(wf)
	if err != nil {
		cm.components.CLogger.Error( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"Impossible to determine the first job in this workflow",
			"id", wf.ID,
			"description", wf.Description,
			"origin", wf.Origin,
			"reason", err.Error(),
		)

		return err
	}

	if err := cm.components.CBackend.InsertWorkflow(wf, job); err != nil {
		cm.components.CLogger.Error( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"Impossible to insert a new workflow",
			"id", wf.ID,
			"description", wf.Description,
			"origin", wf.Origin,
			"reason", err.Error(),
		)

		return err
	}

	cm.newWorkflow(wf)
	cm.newJob(job)

	return nil
}

func (cm *Model) nextStep(job *jw.Job, wf *jw.Workflow) (string, error) {
	step, ok := wf.AllSteps[job.Name]
	if !ok {
		return "", failure.New(nil).
			Set("step", job.Name).
			Msg("this step does not exist") ////////////////////////////////////////////////////////////////////////////
	}

	if step.Next == nil {
		return "", nil
	}

	// AFINIR

	return "unknown", nil
}

func (cm *Model) workflowFinished(job *jw.Job, wf *jw.Workflow) error {
	if *job.WorkflowFailed {
		wf.Status = jw.StatusFailed
	} else {
		wf.Status = jw.StatusSucceeded
	}

	now := time.Now()
	wf.FinishedAt = &now

	if err := cm.components.CBackend.UpdateWorkflow(wf); err != nil {
		return err
	}

	cm.components.CLogger.Info( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
		"Workflow finished",
		"id", wf.ID,
		"description", wf.Description,
		"status", wf.Status,
	)

	cm.publish("workflow.finished", *wf) //*****************************************************************************

	return nil
}

func (m *Model) nextJob(wf *jw.Workflow, pJob *jw.Job, stepName string) (*jw.Job, error) {
	// AFINIR

	return nil, nil
}

func (cm *Model) updateWorkflow(job *jw.Job, wf *jw.Workflow) error {
	step, err := cm.nextStep(job, wf)
	if err != nil {
		return err
	}

	if step == "" {
		return cm.workflowFinished(job, wf)
	}

	job, err = cm.nextJob(wf, job, step)
	if err != nil {
		return err
	}

	return cm.InsertJob(job)
}

/*
######################################################################################################## @(°_°)@ #######
*/
