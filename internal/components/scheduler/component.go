/*
------------------------------------------------------------------------------------------------------------------------
####### scheduler ####### (c) 2020-2021 mls-361 #################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package scheduler

import (
	"github.com/mls-361/minikit"

	"github.com/mls-361/armen/internal/components"
)

type (
	// Scheduler AFAIRE.
	Scheduler struct {
		*minikit.Base
		scheduler *cScheduler
	}
)

// New AFAIRE.
func New(components *components.Components) *Scheduler {
	scheduler := newCScheduler(components)
	components.Scheduler = scheduler

	return &Scheduler{
		Base:      minikit.NewBase("scheduler", "scheduler"),
		scheduler: scheduler,
	}
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
	if err := cs.scheduler.build(); err != nil {
		return err
	}

	cs.Built()

	return nil
}

// Close AFAIRE.
func (cs *Scheduler) Close() {
	cs.scheduler.close()
}

/*
######################################################################################################## @(°_°)@ #######
*/
