/*
------------------------------------------------------------------------------------------------------------------------
####### runner ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package runner

import (
	"time"

	"github.com/mls-361/armen-sdk/jw"
	"github.com/mls-361/armen-sdk/message"
	"github.com/mls-361/failure"

	"github.com/mls-361/armen/internal/components"
)

const (
	_maxAttempts  = 3
	_unknownError = "UNKNOWN ERROR"
)

type (
	// Runner AFAIRE
	Runner struct {
		job        *jw.Job
		components *components.Components
		busCh      chan<- *message.Message
		logger     components.Logger
	}
)

func createLogger(job *jw.Job, components *components.Components) components.Logger {
	if job.Workflow == nil {
		return components.CLogger.CreateLogger(job.ID, "job")
	}

	return components.CLogger.CreateLogger(*job.Workflow, "workflow")
}

// New AFAIRE.
func New(job *jw.Job, components *components.Components, busCh chan<- *message.Message) *Runner {
	return &Runner{
		job:        job,
		components: components,
		busCh:      busCh,
		logger:     createLogger(job, components),
	}
}

func (rr *Runner) publish(topic string) {
	rr.busCh <- message.New(topic, *rr.job)
}

func (rr *Runner) namespaceRunner() (jw.Runner, error) {
	c, err := rr.components.CManager.GetComponent("jw."+rr.job.Namespace, true)
	if err != nil {
		return nil, err
	}

	runner, ok := c.(jw.Runner)
	if !ok {
		return nil, failure.New(nil).
			Set("category", c.Category()).
			Msg("this component is not a job runner") //////////////////////////////////////////////////////////////////
	}

	return runner, nil
}

func (rr *Runner) setError(errMsg string) { rr.job.Error = &errMsg }

func (rr *Runner) removePossibleError() { rr.job.Error = nil }

func (rr *Runner) finished() {
	now := time.Now()
	rr.job.FinishedAt = &now
}

func (rr *Runner) succeeded(jwr *jw.Result) {
	if jwr == nil || jwr.Err == nil {
		rr.removePossibleError()
	} else {
		rr.setError(jwr.Err.Error())
	}

	rr.job.Status = jw.StatusSucceeded

	rr.finished()
}

func (rr *Runner) failed(jwr *jw.Result) {
	var errMsg string

	if jwr.Err == nil {
		errMsg = _unknownError
	} else {
		errMsg = jwr.Err.Error()
	}

	rr.setError(errMsg)

	rr.logger.Error( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
		"The execution of this job has failed",
		"id", rr.job.ID,
		"namespace", rr.job.Namespace,
		"type", rr.job.Type,
		"priority", rr.job.Priority,
		"attempts", rr.job.Attempts,
		"reason", errMsg,
	)

	rr.job.Status = jw.StatusFailed

	rr.finished()
}

func (rr *Runner) pending(jwr *jw.Result) {
	if jwr.Err == nil {
		rr.job.RunAfter = time.Now().Add(jwr.Duration)
	} else {
		rr.job.Attempts++

		if rr.job.Attempts == _maxAttempts {
			rr.failed(jwr)
			return
		}

		errMsg := jwr.Err.Error()

		rr.setError(errMsg)

		rr.logger.Warning( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"The execution of this job has failed",
			"id", rr.job.ID,
			"namespace", rr.job.Namespace,
			"type", rr.job.Type,
			"priority", rr.job.Priority,
			"attempts", rr.job.Attempts,
			"reason", errMsg,
		)

		rr.job.RunAfter = time.Now().Add(time.Duration(rr.job.Attempts) * jwr.Duration)
	}

	rr.job.Status = jw.StatusPending
}

// DoIt AFAIRE.
func (rr *Runner) RunJob() {
	defer rr.logger.Remove()

	if rr.job.Status == jw.StatusToDo {
		rr.logger.Info( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"Begin",
			"name", rr.job.Name,
			"namespace", rr.job.Namespace,
			"type", rr.job.Type,
		)
	} else {
		rr.logger.Info( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"Resume",
			"name", rr.job.Name,
			"namespace", rr.job.Namespace,
			"type", rr.job.Type,
			"attempts", rr.job.Attempts,
		)
	}

	nsRunner, err := rr.namespaceRunner()

	var jwr *jw.Result

	if nsRunner == nil {
		jwr = jw.Failed(err)
	} else {
		rr.job.Status = jw.StatusRunning

		rr.publish("job.run") //****************************************************************************************

		jwr = nsRunner.RunJob(rr.job, rr.logger)
	}

	if jwr == nil {
		rr.succeeded(nil)
	} else {
		switch jwr.Status {
		case jw.StatusSucceeded:
			rr.succeeded(jwr)
		case jw.StatusFailed:
			rr.failed(jwr)
		default:
			rr.pending(jwr)
		}
	}

	if rr.job.Status == jw.StatusPending {
		rr.logger.Info( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"Continuation",
			"after", rr.job.RunAfter.Round(time.Second).String(),
			"attempts", rr.job.Attempts,
		)
	} else {
		rr.logger.Info("End", "status", rr.job.Status) //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	}

	rr.components.CModel.UpdateJob(rr.job)
}

/*
######################################################################################################## @(°_°)@ #######
*/
