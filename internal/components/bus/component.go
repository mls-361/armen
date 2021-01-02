/*
------------------------------------------------------------------------------------------------------------------------
####### bus ####### (c) 2020-2021 mls-361 ########################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package bus

import (
	"github.com/mls-361/minikit"

	"github.com/mls-361/armen/internal/components"
)

type (
	// Bus AFAIRE.
	Bus struct {
		*minikit.Base
		bus *bus
	}
)

// New AFAIRE.
func New(components *components.Components) *Bus {
	bus := newBus(components)
	components.Bus = bus

	return &Bus{
		Base: minikit.NewBase("bus", "bus"),
		bus:  bus,
	}
}

// Dependencies AFAIRE.
func (cb *Bus) Dependencies() []string {
	return []string{
		"application",
		"logger",
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/