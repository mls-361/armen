/*
------------------------------------------------------------------------------------------------------------------------
####### memory ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package memory

import (
	"errors"
	"time"

	"github.com/mls-361/armen-sdk/jw"
)

// InsertJob AFAIRE.
func (cb *Backend) InsertJob(job *jw.Job) (bool, error) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	if job.Key != nil {
		for _, j := range cb.jobs {
			if j.Key != nil && *j.Key == *job.Key &&
				j.Namespace == job.Namespace && j.Type == job.Type &&
				(j.Status == jw.StatusToDo || j.Status == jw.StatusRunning || j.Status == jw.StatusPending) {
				return false, nil
			}
		}
	}

	cb.jobs[job.ID] = job

	return true, nil
}

// NextJob AFAIRE.
func (cb *Backend) NextJob() (*jw.Job, error) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	var job *jw.Job

	now := time.Now()

	for _, j := range cb.jobs {
		if (j.Status != jw.StatusToDo && j.Status != jw.StatusPending) || j.RunAfter.After(now) {
			continue
		}

		if job == nil || j.Priority > job.Priority || j.Weight < job.Weight ||
			j.TimeReference.Before(job.TimeReference) {
			job = j
		}
	}

	if job != nil {
		j := *job

		job.Status = jw.StatusRunning
		job.Weight++

		return &j, nil
	}

	return nil, nil
}

// UpdateJob AFAIRE.
func (cb *Backend) UpdateJob(job *jw.Job) error {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	_, ok := cb.jobs[job.ID]
	if !ok {
		return errors.New("this job does not exist") ///////////////////////////////////////////////////////////////////
	}

	cb.jobs[job.ID] = job

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
