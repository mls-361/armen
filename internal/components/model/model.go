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
	cModel struct {
		components *components.Components
	}
)

func newCModel(components *components.Components) *cModel {
	return &cModel{
		components: components,
	}
}

func (cm *cModel) build() error {
	return nil
}

// AcquireLock AFAIRE.
func (cm *cModel) AcquireLock(name, owner string, duration time.Duration) (bool, error) {
	return cm.components.Backend.AcquireLock(name, owner, duration)
}

// ReleaseLock AFAIRE.
func (cm *cModel) ReleaseLock(name, owner string) error {
	return cm.components.Backend.ReleaseLock(name, owner)
}

/*
######################################################################################################## @(°_°)@ #######
*/
