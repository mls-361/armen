/*
------------------------------------------------------------------------------------------------------------------------
####### workers ####### (c) 2020-2021 mls-361 ###################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package workers

import (
	"github.com/mls-361/minikit"
)

type (
	// Workers AFAIRE.
	Workers struct {
		*minikit.Base
	}
)

// New AFAIRE.
func New() *Workers {
	return &Workers{
		Base: minikit.NewBase("workers", "workers"),
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
func (cw *Workers) Build(_ *minikit.Manager) error {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
