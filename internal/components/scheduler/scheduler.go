/*
------------------------------------------------------------------------------------------------------------------------
####### scheduler ####### (c) 2020-2021 mls-361 #################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package scheduler

import (
	"github.com/mls-361/armen-sdk/message"
	"github.com/mls-361/scheduler"

	"github.com/mls-361/armen/internal/components"
)

type (
	cScheduler struct {
		components *components.Components
		scheduler  *scheduler.Scheduler
		busCh      chan<- *message.Message
	}
)

func newCScheduler(components *components.Components) *cScheduler {
	cs := &cScheduler{
		components: components,
	}

	cs.scheduler = scheduler.New(cs.eventManager)

	return cs
}

func (cs *cScheduler) build() error {
	var events []*scheduler.Event

	if err := cs.components.Config.Decode(&events, false, "components", "scheduler", "events"); err != nil {
		return err
	}

	for _, e := range events {
		cs.components.Logger.Debug( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"Event",
			"name", e.Name,
			"disabled", e.Disabled,
			"after", e.After.String(),
			"repeat", e.Repeat,
		)

		if err := cs.scheduler.AddEvent(e); err != nil {
			return err
		}
	}

	cs.busCh = cs.components.Bus.AddPublisher("scheduler", 1, 1)

	return nil
}

func (cs *cScheduler) eventManager(name string, data interface{}) {
	if cs.components.Leader.AmITheLeader() {
		cs.components.Logger.Info("Emit event", "name", name) //::::::::::::::::::::::::::::::::::::::::::::::::::::::::
		cs.busCh <- message.New(name, data)                   //********************************************************
	}
}

// Start AFAIRE.
func (cs *cScheduler) Start() {
	cs.scheduler.Start()

	cs.components.Logger.Info(">>>Scheduler") //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
}

// Stop AFAIRE.
func (cs *cScheduler) Stop() {
	cs.scheduler.Stop()

	cs.components.Logger.Info("<<<Scheduler") //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
}

func (cs *cScheduler) close() {
	close(cs.busCh)
}

/*
######################################################################################################## @(°_°)@ #######
*/
