/*
------------------------------------------------------------------------------------------------------------------------
####### plugins ####### (c) 2020-2021 mls-361 ###################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package plugins

import (
	"github.com/mls-361/minikit"
)

type (
	// Plugins AFAIRE.
	Plugins struct {
		*minikit.Base
	}
)

// New AFAIRE.
func New() *Plugins {
	return &Plugins{
		Base: minikit.NewBase("plugins", "plugins"),
	}
}

// Dependencies AFAIRE.
func (cp *Plugins) Dependencies() []string {
	return []string{
		"config",
		"logger",
	}
}

// Build AFAIRE.
func (cp *Plugins) Build(_ *minikit.Manager) error {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
