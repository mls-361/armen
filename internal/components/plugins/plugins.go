/*
------------------------------------------------------------------------------------------------------------------------
####### plugins ####### (c) 2020-2021 mls-361 ###################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package plugins

import (
	"github.com/mls-361/component"
)

type (
	// Plugins AFAIRE.
	Plugins struct {
		*component.Base
	}
)

// New AFAIRE.
func New() *Plugins {
	return &Plugins{
		Base: component.NewBase("plugins", "plugins"),
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
func (cp *Plugins) Build(_ *component.Manager) error {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
