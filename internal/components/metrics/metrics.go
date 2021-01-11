/*
------------------------------------------------------------------------------------------------------------------------
####### metrics ####### (c) 2020-2021 mls-361 ###################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package metrics

import (
	"github.com/mls-361/armen-sdk/components"
	"github.com/mls-361/metrics"
	"github.com/mls-361/minikit"
)

type (
	// Metrics AFAIRE.
	Metrics struct {
		*minikit.Base
		metrics.Metrics
		components *components.Components
	}
)

// New AFAIRE.
func New(components *components.Components) *Metrics {
	cm := &Metrics{
		Base:       minikit.NewBase("metrics", "metrics"),
		Metrics:    metrics.New(),
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
	cm.components.Router.Get("/metrics", cm.Handler())
	return nil
}

// NewCounter AFAIRE.

/*
######################################################################################################## @(°_°)@ #######
*/
