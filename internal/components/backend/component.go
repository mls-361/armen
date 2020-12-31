/*
------------------------------------------------------------------------------------------------------------------------
####### backend ####### (c) 2020-2021 mls-361 ###################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package backend

import (
	"github.com/mls-361/armen/internal/components"
	"github.com/mls-361/component"
)

type (
	// Backend AFAIRE.
	Backend struct {
		*component.Base
	}
)

// New AFAIRE.
func New(components *components.Components) *Backend {
	return &Backend{
		Base: component.NewBase("backend", "backend"),
	}
}

// Dependencies AFAIRE.
func (cb *Backend) Dependencies() []string {
	return []string{
		"config",
		"logger",
	}
}

// Build AFAIRE.
func (cb *Backend) Build(_ *component.Manager) error {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
