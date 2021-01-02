/*
------------------------------------------------------------------------------------------------------------------------
####### leader ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package leader

import (
	"github.com/mls-361/minikit"

	"github.com/mls-361/armen/internal/components"
)

type (
	// Leader AFAIRE.
	Leader struct {
		*minikit.Base
		leader *leader
	}
)

// New AFAIRE.
func New(components *components.Components) *Leader {
	leader := newLeader(components)
	components.Leader = leader

	return &Leader{
		Base:   minikit.NewBase("leader", "leader"),
		leader: leader,
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
	if err := cl.leader.build(); err != nil {
		return err
	}

	cl.Built()

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
