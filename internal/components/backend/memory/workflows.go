/*
------------------------------------------------------------------------------------------------------------------------
####### memory ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package memory

import (
	"errors"

	"github.com/mls-361/armen-sdk/jw"
)

// Workflow AFAIRE.
func (cb *Backend) Workflow(id string, mustExist bool) (*jw.Workflow, error) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	wf, ok := cb.workflows[id]
	if !ok {
		if mustExist {
			return nil, errors.New("this workflow does not exist") /////////////////////////////////////////////////////
		}

		return nil, nil
	}

	return wf, nil
}

// InsertWorkflow AFAIRE.
func (cb *Backend) InsertWorkflow(wf *jw.Workflow, job *jw.Job) error {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	cb.workflows[wf.ID] = wf
	cb.jobs[job.ID] = job

	return nil
}

// UpdateWorkflow AFAIRE.
func (cb *Backend) UpdateWorkflow(wf *jw.Workflow) error {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	cb.workflows[wf.ID] = wf

	return nil
}

func (cb *Backend) deleteFinishedWorkflows() int {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	count := 0

	for id, wf := range cb.workflows {
		if wf.Status == jw.StatusFailed || wf.Status == jw.StatusSucceeded {
			delete(cb.workflows, id)
			count++
		}
	}

	return count
}

/*
######################################################################################################## @(°_°)@ #######
*/
