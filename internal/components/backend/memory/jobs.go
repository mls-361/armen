/*
------------------------------------------------------------------------------------------------------------------------
####### memory ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package memory

import "github.com/mls-361/armen-sdk/jw"

// InsertJob AFAIRE.
func (cb *Backend) InsertJob(job *jw.Job) (bool, error) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	if job.UniqueKey != nil {
		for _, j := range cb.jobs {
			if j.UniqueKey != nil && *j.UniqueKey == *job.UniqueKey &&
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
func (cb *Backend) NextJob() *jw.Job {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
