/*
------------------------------------------------------------------------------------------------------------------------
####### workers ####### (c) 2020-2021 mls-361 ###################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package workers

import (
	"github.com/mls-361/component"
)

type (
	// Workers AFAIRE.
	Workers struct {
		*component.Base
	}
)

// New AFAIRE.
func New() *Workers {
	return &Workers{
		Base: component.NewBase("workers", "workers"),
	}
}

// Dependencies AFAIRE.
func (cw *Workers) Dependencies() []string {
	return []string{
		"bus",
		"config",
		"logger",
	}
}

// Build AFAIRE.
func (cw *Workers) Build(_ *component.Manager) error {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
