/*
------------------------------------------------------------------------------------------------------------------------
####### pgsql ####### (c) 2020-2021 mls-361 ######################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package pgsql

import (
	"context"
	"time"

	"github.com/mls-361/armen-sdk/jw"
	"github.com/mls-361/pgsql"
)

func (cb *Backend) addJobToHistory(t *pgsql.Transaction, status string, job *jw.Job) error {
	_, err := t.Execute(
		"INSERT INTO history (created_at, job, workflow, type, status, data) VALUES ($1, $2, $3, $4, $5, $6)",
		time.Now(),
		job.ID,
		job.Workflow,
		job.Type,
		status,
		job,
	)
	return err
}

func (cb *Backend) addWorkflowToHistory(t *pgsql.Transaction, status string, wf *jw.Workflow) error {
	_, err := t.Execute(
		"INSERT INTO history (created_at, workflow, type, status, data) VALUES ($1, $2, $3, $4, $5)",
		time.Now(),
		wf.ID,
		wf.Type,
		status,
		wf,
	)
	return err
}

func (cb *Backend) deleteOldestHistory(ctx context.Context, client *pgsql.Client) error {
	_, err := client.Execute(ctx, "DELETE FROM history WHERE created_at <= $1", time.Now().Add(cb.historyRT))
	return err
}

/*
######################################################################################################## @(°_°)@ #######
*/
