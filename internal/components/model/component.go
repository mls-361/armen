/*
------------------------------------------------------------------------------------------------------------------------
####### model ####### (c) 2020-2021 mls-361 ######################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package model

import (
	"github.com/mls-361/minikit"
)

type (
	// Model AFAIRE.
	Model struct {
		*minikit.Base
	}
)

// New AFAIRE.
func New() *Model {
	return &Model{
		Base: minikit.NewBase("model", "model"),
	}
}

// Dependencies AFAIRE.
func (cm *Model) Dependencies() []string {
	return []string{
		"backend",
		"config",
		"crypto",
		"logger",
	}
}

// Build AFAIRE.
func (cm *Model) Build(_ *minikit.Manager) error {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
