/*
------------------------------------------------------------------------------------------------------------------------
####### api ####### (c) 2020-2021 mls-361 ########################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package api

import (
	"github.com/mls-361/armen/internal/components"
)

type (
	API struct {
		components *components.Components
	}
)

// New AFAIRE.
func New(cs *components.Components) *API {
	return &API{
		components: cs,
	}
}

// Setup AFAIRE.
func (api *API) Setup() {
}

// AppStopping AFAIRE.
func (api *API) AppStopping() {}

/*
######################################################################################################## @(°_°)@ #######
*/
