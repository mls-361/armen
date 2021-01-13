/*
------------------------------------------------------------------------------------------------------------------------
####### pgsql ####### (c) 2020-2021 mls-361 ######################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package pgsql

import "github.com/mls-361/armen/internal/components"

type (
	// Backend AFAIRE.
	Backend struct {
		components *components.Components
	}
)

func New(components *components.Components) *Backend {
	return &Backend{
		components: components,
	}
}

// Build AFAIRE.
func (cb *Backend) Build() error {
	return nil
}

// Close AFAIRE.
func (cb *Backend) Close() {}

/*
######################################################################################################## @(°_°)@ #######
*/
