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

func (cb *Backend) tryInsertJob(t *pgsql.Transaction, job *jw.Job) error {
	_, err := t.Execute(
		`
		INSERT INTO jobs (id, name, namespace, type, origin, priority, key, workflow,
        workflow_failed, emails, config, private, public, created_at, status, error,
		attempts, finished_at, run_after, result, next_step, weight, time_reference)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15,
		$16, $17, $18, $19, $20, $21, $22, $23)`,
		job.ID,
		job.Name,
		job.Namespace,
		job.Type,
		job.Origin,
		job.Priority,
		job.Key,
		job.Workflow,
		job.WorkflowFailed,
		job.Emails,
		job.Config,
		job.Private,
		job.Public,
		job.CreatedAt,
		job.Status,
		job.Error,
		job.Attempts,
		job.FinishedAt,
		job.RunAfter,
		job.Result,
		job.NextStep,
		job.Weight,
		job.TimeReference,
	)
	return err
}

func (cb *Backend) existJob(t *pgsql.Transaction, job *jw.Job) (bool, error) {
	var id string

	err := t.QueryRow(
		`
		SELECT id
		FROM jobs
		WHERE namespace = $1 AND type = $2 AND key = $3 AND (status = $4 OR status = $5 OR status = $6)
		LIMIT 1
		FOR UPDATE`,
		job.Namespace,
		job.Type,
		job.Key,
		jw.StatusToDo,
		jw.StatusRunning,
		jw.StatusPending,
	).Scan(&id)

	if err == nil {
		return true, nil
	}

	if errors.Is(err, pgsql.ErrNoRows) {
		return false, nil
	}

	return false, err
}

func (cb *Backend) maybeInsertJob(job *jw.Job) (bool, error) {
	client, err := cb.primary()
	if err != nil {
		return false, err
	}

	ctx, cancel := client.ContextWT(5 * time.Second)
	defer cancel()

	inserted := false

	err = client.Transaction(
		ctx,
		func(t *pgsql.Transaction) error {
			if exist, err := cb.existJob(t, job); exist || err != nil {
				return err
			}

			if err := cb.tryInsertJob(t, job); err != nil {
				return err
			}

			if err := cb.addJobToHistory(t, "add", job); err != nil {
				return err
			}

			inserted = true

			return nil
		},
	)

	return inserted, err
}

func (cb *Backend) insertJob(job *jw.Job) error {
	client, err := cb.primary()
	if err != nil {
		return err
	}

	ctx, cancel := client.ContextWT(5 * time.Second)
	defer cancel()

	return client.Transaction(
		ctx,
		func(t *pgsql.Transaction) error {
			if err := cb.tryInsertJob(t, job); err != nil {
				return err

			}

			return cb.addJobToHistory(t, "add", job)
		},
	)
}

// InsertJob AFAIRE.
func (cb *Backend) InsertJob(job *jw.Job) (bool, error) {
	if job.Key != nil && *job.Key != "" {
		return cb.maybeInsertJob(job)
	}

	return true, cb.insertJob(job)
}

// NextJob AFAIRE.
func (cb *Backend) NextJob() (*jw.Job, error) {
	return nil, nil
}

// UpdateJob AFAIRE.
func (cb *Backend) UpdateJob(job *jw.Job) error {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
