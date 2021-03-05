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

	workflowFailed := false

	job.Public = wf.Data
	job.WorkflowFailed = &workflowFailed

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

func (cm *Model) nextLabelStep(job *jw.Job, label string, value interface{}) (string, map[string]interface{}, error) {
	switch next := value.(type) {
	case nil:
		return "", nil, nil
	case string:
		return next, nil, nil
	case map[string]interface{}:
		vs, ok := next["step"]
		if !ok {
			return "", nil, failure.New(nil).
				Set("job", job.Name).
				Set("label", label).
				Msg("the 'step' key is missing") ///////////////////////////////////////////////////////////////////////
		}

		switch name := vs.(type) {
		case nil:
			return "", nil, nil
		case string:
			vd, ok := next["data"]
			if !ok {
				return name, nil, nil
			}

			switch data := vd.(type) {
			case nil:
				return name, nil, nil
			case map[string]interface{}:
				return name, data, nil
			default:
				return "", nil, failure.New(nil).
					Set("job", job.Name).
					Set("label", label).
					Msg("the 'data' key is not valid") /////////////////////////////////////////////////////////////////
			}
		default:
			return "", nil, failure.New(nil).
				Set("job", job.Name).
				Set("label", label).
				Msg("the 'step' key is not valid") /////////////////////////////////////////////////////////////////////
		}
	default:
		return "", nil, failure.New(nil).
			Set("job", job.Name).
			Set("label", label).
			Msg("the configuration for this label is not valid") ///////////////////////////////////////////////////////
	}
}

func (cm *Model) nextStep(job *jw.Job, wf *jw.Workflow) (string, map[string]interface{}, error) {
	if job.NextStep != nil {
		return *job.NextStep, nil, nil
	}

	step, ok := wf.AllSteps[job.Name]
	if !ok {
		return "", nil, failure.New(nil).
			Set("step", job.Name).
			Msg("strangely, this step does not seem to exist") /////////////////////////////////////////////////////////
	}

	if step.Next == nil {
		return "", nil, nil
	}

	var result string

	if job.Result != nil {
		result = *job.Result

		value, ok := step.Next[result]
		if ok {
			return cm.nextLabelStep(job, result, value)
		}
	}

	status := job.Status.String()

	value, ok := step.Next[status]
	if ok {
		return cm.nextLabelStep(job, status, value)
	}

	f := failure.New(nil).Set("job", job.Name)

	if job.Result != nil {
		_ = f.Set("result", result)
	}

	return "", nil, f.Set("status", status).Msg("impossible to determine the next step") ///////////////////////////////
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

func (m *Model) nextJob(wf *jw.Workflow, pj *jw.Job, stepName string, data map[string]interface{}) (*jw.Job, error) {
	job, err := m.stepToJob(wf, stepName)
	if err != nil {
		return nil, err
	}

	job.Private = data
	job.Public = pj.Public
	job.WorkflowFailed = pj.WorkflowFailed

	return job, nil
}

func (cm *Model) updateWorkflow(job *jw.Job, wf *jw.Workflow) error {
	stepName, data, err := cm.nextStep(job, wf)
	if err != nil {
		return err
	}

	if stepName == "" {
		return cm.workflowFinished(job, wf)
	}

	job, err = cm.nextJob(wf, job, stepName, data)
	if err != nil {
		return err
	}

	return cm.InsertJob(job)
}

func (cm *Model) deleteFinishedWorkflows() {
	count, err := cm.components.CBackend.DeleteFinishedWorkflows()
	if err != nil {
		cm.components.CLogger.Error("Impossible to delete finished workflows") //:::::::::::::::::::::::::::::::::::::::
		return
	}

	cm.components.CLogger.Info("Workflows deleted", "count", count) //::::::::::::::::::::::::::::::::::::::::::::::::::
}

/*
######################################################################################################## @(°_°)@ #######
*/
