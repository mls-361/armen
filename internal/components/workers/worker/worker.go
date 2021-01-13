/*
------------------------------------------------------------------------------------------------------------------------
####### worker ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package worker

import (
	"time"

	"github.com/mls-361/armen-sdk/message"
	"github.com/mls-361/armen-sdk/worker"

	"github.com/mls-361/armen/internal/components"
	"github.com/mls-361/armen/internal/components/workers/runner"
)

type (
	// Worker AFAIRE.
	Worker struct {
		*worker.Worker
		components *components.Components
		busCh      chan<- *message.Message
		stopCh     chan struct{}
		logger     components.Logger
	}
)

// New AFAIRE.
func New(components *components.Components, busCh chan<- *message.Message, stopCh chan struct{}) *Worker {
	worker := worker.New()

	return &Worker{
		Worker:     worker,
		components: components,
		busCh:      busCh,
		stopCh:     stopCh,
		logger:     components.CLogger.CreateLogger(worker.ID, "worker"),
	}
}

func (w *Worker) publish(topic string, data interface{}) {
	w.Data = data
	w.busCh <- message.New(topic, *w.Worker)
	w.Data = nil
}

func (w *Worker) maybeRunJob() time.Duration {
	job := w.components.CModel.NextJob()

	if job == nil {
		return 1
	}

	w.Jobs++

	w.logger.Info("Run job", "id", job.ID, "count", w.Jobs) //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

	w.publish("worker.busy", *job) //***********************************************************************************

	runner.New(job, w.components, w.busCh).DoIt()

	w.publish("worker.free", nil) //************************************************************************************

	return 0
}

func (w *Worker) run() {
	for {
		select {
		case <-w.stopCh:
			return
		case <-time.After(w.maybeRunJob() * time.Second):
		}
	}
}

// Run AFAIRE.
func (w *Worker) Run() {
	w.logger.Info(">>>Worker", "id", w.ID) //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

	w.publish("worker.start", nil) //***********************************************************************************

	w.run()

	w.publish("worker.stop", nil) //************************************************************************************

	w.logger.Info("<<<Worker") //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

	w.logger.RemoveLogger("")
}

/*
######################################################################################################## @(°_°)@ #######
*/
