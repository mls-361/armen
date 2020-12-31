/*
------------------------------------------------------------------------------------------------------------------------
####### scheduler ####### (c) 2020-2021 mls-361 #################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package scheduler

import (
	"github.com/mls-361/component"
)

type (
	// Scheduler AFAIRE.
	Scheduler struct {
		*component.Base
	}
)

// New AFAIRE.
func New() *Scheduler {
	return &Scheduler{
		Base: component.NewBase("scheduler", "scheduler"),
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
func (cs *Scheduler) Build(_ *component.Manager) error {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
