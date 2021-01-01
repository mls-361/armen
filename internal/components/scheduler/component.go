/*
------------------------------------------------------------------------------------------------------------------------
####### scheduler ####### (c) 2020-2021 mls-361 #################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package scheduler

import (
	"github.com/mls-361/minikit"
)

type (
	// Scheduler AFAIRE.
	Scheduler struct {
		*minikit.Base
	}
)

// New AFAIRE.
func New() *Scheduler {
	return &Scheduler{
		Base: minikit.NewBase("scheduler", "scheduler"),
	}
}

// Dependencies AFAIRE.
func (cs *Scheduler) Dependencies() []string {
	return []string{
		"bus",
		"config",
		"logger",
	}
}

// Build AFAIRE.
func (cs *Scheduler) Build(_ *minikit.Manager) error {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
