/*
------------------------------------------------------------------------------------------------------------------------
####### leader ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package leader

import (
	"github.com/mls-361/component"
)

type (
	// Leader AFAIRE.
	Leader struct {
		*component.Base
	}
)

// New AFAIRE.
func New() *Leader {
	return &Leader{
		Base: component.NewBase("leader", "leader"),
	}
}

// Dependencies AFAIRE.
func (cl *Leader) Dependencies() []string {
	return []string{
		"logger",
		"model",
	}
}

// Build AFAIRE.
func (cl *Leader) Build(_ *component.Manager) error {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
