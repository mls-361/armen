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
	"github.com/mls-361/failure"
	"github.com/mls-361/pgsql"
)

func (cb *Backend) tryInsertJob(t *pgsql.Transaction, job *jw.Job) error {
	_, err := t.Execute(
		`
		INSERT INTO jobs
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12,
		$13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)`,
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
		LIMIT 1`,
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

	ctx, cancel := client.ContextWT(10 * time.Second)
	defer cancel()

	inserted := false

	err = client.Transaction(
		ctx,
		func(t *pgsql.Transaction) error {
			if err := cb.advisoryLock(t, _lockInsertJob); err != nil {
				return err
			}

			if exist, err := cb.existJob(t, job); exist || err != nil {
				return err
			}

			if err := cb.tryInsertJob(t, job); err != nil {
				return err
			}

			if err := cb.addJobToHistory(t, "insert", job); err != nil {
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

			return cb.addJobToHistory(t, "insert", job)
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
	client, err := cb.primary()
	if err != nil {
		return nil, err
	}

	ctx, cancel := client.ContextWT(5 * time.Second)
	defer cancel()

	var job jw.Job

	if err := client.Transaction(
		ctx,
		func(t *pgsql.Transaction) error {
			if err := t.QueryRow(
				`
				SELECT *
				FROM jobs
				WHERE (status = $1 OR status = $2) AND run_after <= $3
				ORDER BY priority DESC, weight ASC, time_reference ASC
				LIMIT 1
				FOR UPDATE`,
				jw.StatusToDo,
				jw.StatusPending,
				time.Now(),
			).Scan(
				&job.ID,
				&job.Name,
				&job.Namespace,
				&job.Type,
				&job.Origin,
				&job.Priority,
				&job.Key,
				&job.Workflow,
				&job.WorkflowFailed,
				&job.Emails,
				&job.Config,
				&job.Private,
				&job.Public,
				&job.CreatedAt,
				&job.Status,
				&job.Error,
				&job.Attempts,
				&job.FinishedAt,
				&job.RunAfter,
				&job.Result,
				&job.NextStep,
				&job.Weight,
				&job.TimeReference,
			); err != nil {
				return err
			}

			job.Weight++

			count, err := t.Execute(
				"UPDATE jobs SET status = $1, weight = $2 WHERE id = $3",
				jw.StatusRunning,
				job.Weight,
				job.ID,
			)
			if err != nil {
				return err
			}

			if count != 1 {
				return failure.New(nil).
					Set("job", job.ID).
					Msg("impossible to update this job") ///////////////////////////////////////////////////////////////
			}

			return cb.addJobToHistory(t, "select", &job)
		},
	); err != nil {
		if errors.Is(err, pgsql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &job, nil
}

// UpdateJob AFAIRE.
func (cb *Backend) UpdateJob(job *jw.Job) error {
	client, err := cb.primary()
	if err != nil {
		return err
	}

	ctx, cancel := client.ContextWT(5 * time.Second)
	defer cancel()

	return client.Transaction(
		ctx,
		func(t *pgsql.Transaction) error {
			count, err := t.Execute(
				`
				UPDATE jobs
				SET workflow_failed = $1,private = $2, public = $3, status = $4,
				error = $5, attempts = $6, finished_at = $7, run_after = $8, result = $9,
				next_step = $10
				WHERE id = $11`,
				job.WorkflowFailed,
				job.Private,
				job.Public,
				job.Status,
				job.Error,
				job.Attempts,
				job.FinishedAt,
				job.RunAfter,
				job.Result,
				job.NextStep,
				job.ID,
			)
			if err != nil {
				return err
			}

			if count != 1 {
				return failure.New(nil).
					Set("job", job.ID).
					Msg("impossible to update this job") ///////////////////////////////////////////////////////////////
			}

			return cb.addJobToHistory(t, "update", job)
		},
	)
}

// DeleteFinishedJobs AFAIRE.
func (cb *Backend) DeleteFinishedJobs() (int64, error) {
	client, err := cb.primary()
	if err != nil {
		return 0, err
	}

	ctx, cancel := client.ContextWT(5 * time.Second)
	defer cancel()

	return client.Execute(
		ctx,
		"DELETE FROM jobs WHERE workflow IS NULL AND (status = $1 OR status = $2)",
		jw.StatusFailed,
		jw.StatusSucceeded,
	)
}

/*
######################################################################################################## @(°_°)@ #######
*/
