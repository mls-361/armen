/*
------------------------------------------------------------------------------------------------------------------------
####### runner ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package runner

import (
	"github.com/mls-361/armen-sdk/jw"
	"github.com/mls-361/armen-sdk/message"

	"github.com/mls-361/armen/internal/components"
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

// DoIt AFAIRE.
func (rr *Runner) DoIt() {
	defer rr.logger.RemoveLogger("")
}

/*
######################################################################################################## @(°_°)@ #######
*/
