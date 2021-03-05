/*
------------------------------------------------------------------------------------------------------------------------
####### pgsql ####### (c) 2020-2021 mls-361 ######################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package pgsql

import (
	"errors"
	"time"

	"github.com/mls-361/armen-sdk/jw"
	"github.com/mls-361/pgsql"
)

// Workflow AFAIRE.
func (cb *Backend) Workflow(id string, mustExist bool) (*jw.Workflow, error) {
	client, err := cb.primaryPreferred()
	if err != nil {
		return nil, err
	}

	ctx, cancel := client.ContextWT(5 * time.Second)
	defer cancel()

	var wf jw.Workflow

	if err := client.QueryRow(ctx, "SELECT * FROM workflows WHERE id = $1", id).Scan(
		&wf.ID,
		&wf.Description,
		&wf.Origin,
		&wf.Priority,
		&wf.FirstStep,
		&wf.AllSteps,
		&wf.ExternalReference,
		&wf.Emails,
		&wf.Data,
		&wf.CreatedAt,
		&wf.Status,
		&wf.FinishedAt,
	); err != nil {
		if errors.Is(err, pgsql.ErrNoRows) {
			if mustExist {
				return nil, errors.New("this workflow does not exist") /////////////////////////////////////////////////
			}

			return nil, nil
		}

		return nil, err
	}

	return &wf, nil
}

// InsertWorkflow AFAIRE.
func (cb *Backend) InsertWorkflow(wf *jw.Workflow, job *jw.Job) error {
	client, err := cb.primary()
	if err != nil {
		return err
	}

	ctx, cancel := client.ContextWT(10 * time.Second)
	defer cancel()

	return client.Transaction(
		ctx,
		func(t *pgsql.Transaction) error {
			_, err := t.Execute(
				"INSERT INTO workflows VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)",
				wf.ID,
				wf.Description,
				wf.Origin,
				wf.Priority,
				wf.FirstStep,
				wf.AllSteps,
				wf.ExternalReference,
				wf.Emails,
				wf.Data,
				wf.CreatedAt,
				wf.Status,
				wf.FinishedAt,
			)

			if err != nil {
				return err
			}

			if err := cb.addWorkflowToHistory(t, "insert", wf); err != nil {
				return err
			}

			if err := cb.tryInsertJob(t, job); err != nil {
				return err
			}

			return cb.addJobToHistory(t, "insert", job)
		},
	)
}

// UpdateWorkflow AFAIRE.
func (cb *Backend) UpdateWorkflow(wf *jw.Workflow) error {
	client, err := cb.primary()
	if err != nil {
		return err
	}

	ctx, cancel := client.ContextWT(5 * time.Second)
	defer cancel()

	return client.Transaction(
		ctx,
		func(t *pgsql.Transaction) error {
			_, err := t.Execute(
				"UPDATE workflows SET status = $1, finished_at = $2 WHERE id = $3",
				wf.Status,
				wf.FinishedAt,
				wf.ID,
			)

			if err != nil {
				return err
			}

			return cb.addWorkflowToHistory(t, "update", wf)
		},
	)
}

// DeleteFinishedWorkflows AFAIRE.
func (cb *Backend) DeleteFinishedWorkflows() (int64, error) {
	client, err := cb.primary()
	if err != nil {
		return 0, err
	}

	ctx, cancel := client.ContextWT(10 * time.Second)
	defer cancel()

	return client.Execute(
		ctx,
		"DELETE FROM workflows WHERE (status = $1 OR status = $2)",
		jw.StatusFailed,
		jw.StatusSucceeded,
	)
}

/*
######################################################################################################## @(°_°)@ #######
*/
