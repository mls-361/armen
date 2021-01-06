/*
------------------------------------------------------------------------------------------------------------------------
####### pgsql ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package pgsql

import (
	"time"

	"github.com/mls-361/armen-sdk/components"
)

type (
	// Backend AFAIRE.
	Backend struct {
		components *components.Components
	}
)

func New(components *components.Components) *Backend {
	return &Backend{
		components: components,
	}
}

// Build AFAIRE.
func (cb *Backend) Build() error {
	return nil
}

// AcquireLock AFAIRE.
func (cb *Backend) AcquireLock(name, owner string, duration time.Duration) (bool, error) {
	return true, nil
}

// ReleaseLock AFAIRE.
func (cb *Backend) ReleaseLock(name, owner string) error {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
