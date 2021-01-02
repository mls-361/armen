/*
------------------------------------------------------------------------------------------------------------------------
####### leader ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package leader

import "github.com/mls-361/armen/internal/components"

type (
	leader struct {
		components *components.Components
	}
)

func newLeader(components *components.Components) *leader {
	return &leader{
		components: components,
	}
}

func (cl *leader) build() error {
	return nil
}

func (cl *leader) AmITheLeader() bool {
	return false
}

/*
######################################################################################################## @(°_°)@ #######
*/
