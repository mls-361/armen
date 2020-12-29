/*
------------------------------------------------------------------------------------------------------------------------
####### bus ####### (c) 2020-2021 mls-361 ########################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package bus

import (
	"github.com/mls-361/component"
)

type (
	// Bus AFAIRE.
	Bus struct {
		*component.Base
	}
)

// New AFAIRE.
func New() *Bus {
	return &Bus{
		Base: component.NewBase("bus", "bus"),
	}
}

// Dependencies AFAIRE.
func (cb *Bus) Dependencies() []string {
	return []string{
		"logger",
	}
}

// Build AFAIRE.
func (cb *Bus) Build(_ *component.Manager) error {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
