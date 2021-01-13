/*
------------------------------------------------------------------------------------------------------------------------
####### metrics ####### (c) 2020-2021 mls-361 ###################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package metrics

import (
	"github.com/mls-361/metrics"
	"github.com/mls-361/minikit"

	"github.com/mls-361/armen/internal/components"
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

	components.CMetrics = cm

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
	cm.components.CRouter.Get("/metrics", cm.Handler())
	return nil
}

// NewCounter AFAIRE.

/*
######################################################################################################## @(°_°)@ #######
*/
