/*
------------------------------------------------------------------------------------------------------------------------
####### model ####### (c) 2020-2021 mls-361 ######################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package model

import (
	"github.com/mls-361/armen-sdk/components"
	"github.com/mls-361/minikit"
)

type (
	// Model AFAIRE.
	Model struct {
		*minikit.Base
		model *cModel
	}
)

// New AFAIRE.
func New(components *components.Components) *Model {
	model := newCModel(components)
	components.Model = model

	return &Model{
		Base:  minikit.NewBase("model", "model"),
		model: model,
	}
}

// Dependencies AFAIRE.
func (cm *Model) Dependencies() []string {
	return []string{
		"backend",
		"logger",
	}
}

// Build AFAIRE.
func (cm *Model) Build(_ *minikit.Manager) error {
	if err := cm.model.build(); err != nil {
		return err
	}

	cm.Built()

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
