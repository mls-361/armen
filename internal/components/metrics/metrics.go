/*
------------------------------------------------------------------------------------------------------------------------
####### metrics ####### (c) 2020-2021 mls-361 ###################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package metrics

import (
	"github.com/mls-361/armen-sdk/components"
	"github.com/mls-361/minikit"
)

type (
	// Metrics AFAIRE.
	Metrics struct {
		*minikit.Base
		*tree
		components *components.Components
	}
)

// New AFAIRE.
func New(components *components.Components) *Metrics {
	cm := &Metrics{
		Base:       minikit.NewBase("metrics", "metrics"),
		tree:       newTree(),
		components: components,
	}

	components.Metrics = cm

	return cm
}

// Dependencies AFAIRE.
func (cm *Metrics) Dependencies() []string {
	return []string{
		"router",
	}
}

// Build AFAIRE.
func (cm *Metrics) Build(_ *minikit.Manager) error {
	cm.components.Router.Get("/metrics", cm.handler())

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
