/*
------------------------------------------------------------------------------------------------------------------------
####### pgsql ####### (c) 2020-2021 mls-361 ######################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package pgsql

import (
	"time"

	"github.com/mls-361/armen-sdk/jw"
	"github.com/mls-361/pgsql"
)

func (cb *Backend) addJobToHistory(t *pgsql.Transaction, action string, job *jw.Job) error {
	_, err := t.Execute(
		"INSERT INTO history (created_at, action, job, workflow, data) VALUES ($1, $2, $3, $4, $5)",
		time.Now(),
		action,
		job.ID,
		job.Workflow,
		job,
	)
	return err
}

func (cb *Backend) addWorkflowToHistory(t *pgsql.Transaction, action string, wf *jw.Workflow) error {
	_, err := t.Execute(
		"INSERT INTO history (created_at, action, workflow, data) VALUES ($1, $2, $3, $4)",
		time.Now(),
		action,
		wf.ID,
		wf,
	)
	return err
}

/*
######################################################################################################## @(°_°)@ #######
*/
