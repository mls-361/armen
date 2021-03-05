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

const (
	_retentionTime = 7 * 24 * time.Hour // AFINIR: configurable
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

func (cb *Backend) deleteOldestHistory(ctx context.Context, client *pgsql.Client) error {
	_, err := client.Execute(ctx, "DELETE FROM history WHERE created_at <= $1", time.Now().Add(-1*_retentionTime))
	return err
}

/*
######################################################################################################## @(°_°)@ #######
*/
