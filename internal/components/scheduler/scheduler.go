/*
------------------------------------------------------------------------------------------------------------------------
####### scheduler ####### (c) 2020-2021 mls-361 #################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package scheduler

import (
	"github.com/mls-361/armen-sdk/message"
	"github.com/mls-361/minikit"
	"github.com/mls-361/scheduler"

	"github.com/mls-361/armen/internal/components"
)

type (
	// Scheduler AFAIRE.
	Scheduler struct {
		*minikit.Base
		components *components.Components
		scheduler  *scheduler.Scheduler
		busCh      chan<- *message.Message
	}
)

// New AFAIRE.
func New(components *components.Components) *Scheduler {
	cs := &Scheduler{
		Base:       minikit.NewBase("scheduler", ""),
		components: components,
	}

	cs.scheduler = scheduler.New(cs.eventManager)

	components.CScheduler = cs

	return cs
}

// Dependencies AFAIRE.
func (cs *Scheduler) Dependencies() []string {
	return []string{
		"bus",
		"config",
		"leader",
		"logger",
	}
}

// Build AFAIRE.
func (cs *Scheduler) Build(_ *minikit.Manager) error {
	var events []*scheduler.Event

	if err := cs.components.CConfig.Decode(&events, false, "components", "scheduler", "events"); err != nil {
		return err
	}

	for _, e := range events {
		cs.components.CLogger.Debug( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
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

	cs.busCh = cs.components.CBus.AddPublisher("scheduler", 3, 1)

	return nil
}

func (cs *Scheduler) eventManager(name string, data interface{}) {
	if cs.components.CLeader.AmITheLeader() {
		cs.components.CLogger.Info("Emit event", "name", name) //:::::::::::::::::::::::::::::::::::::::::::::::::::::::
		cs.busCh <- message.New(name, data)                    //*******************************************************
	}
}

// Start AFAIRE.
func (cs *Scheduler) Start() {
	cs.scheduler.Start()

	cs.components.CLogger.Info(">>>Scheduler") //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
}

// Stop AFAIRE.
func (cs *Scheduler) Stop() {
	cs.scheduler.Stop()

	cs.components.CLogger.Info("<<<Scheduler") //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
}

// Close AFAIRE.
func (cs *Scheduler) Close() {
	close(cs.busCh)
}

/*
######################################################################################################## @(°_°)@ #######
*/
