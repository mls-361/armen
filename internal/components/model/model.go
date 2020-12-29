/*
------------------------------------------------------------------------------------------------------------------------
####### model ####### (c) 2020-2021 mls-361 ######################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package model

import (
	"github.com/mls-361/component"
)

type (
	// Model AFAIRE.
	Model struct {
		*component.Base
	}
)

// New AFAIRE.
func New() *Model {
	return &Model{
		Base: component.NewBase("model", "model"),
	}
}

// Dependencies AFAIRE.
func (cm *Model) Dependencies() []string {
	return []string{
		"config",
		"crypto",
	}
}

// Build AFAIRE.
func (cm *Model) Build(_ *component.Manager) error {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
