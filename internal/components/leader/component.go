/*
------------------------------------------------------------------------------------------------------------------------
####### leader ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package leader

import (
	"github.com/mls-361/minikit"
)

type (
	// Leader AFAIRE.
	Leader struct {
		*minikit.Base
	}
)

// New AFAIRE.
func New() *Leader {
	return &Leader{
		Base: minikit.NewBase("leader", "leader"),
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
func (cl *Leader) Build(_ *minikit.Manager) error {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
