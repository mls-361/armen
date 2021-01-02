/*
------------------------------------------------------------------------------------------------------------------------
####### model ####### (c) 2020-2021 mls-361 ######################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package model

import (
	"time"

	"github.com/mls-361/armen/internal/components"
)

type (
	model struct {
		components *components.Components
	}
)

func newModel(components *components.Components) *model {
	return &model{
		components: components,
	}
}

func (cm *model) build() error {
	return nil
}

// AcquireLock AFAIRE.
func (cm *model) AcquireLock(name, owner string, duration time.Duration) (bool, error) {
	return cm.components.Backend.AcquireLock(name, owner, duration)
}

// ReleaseLock AFAIRE.
func (cm *model) ReleaseLock(name, owner string) error {
	return cm.components.Backend.ReleaseLock(name, owner)
}

/*
######################################################################################################## @(°_°)@ #######
*/
