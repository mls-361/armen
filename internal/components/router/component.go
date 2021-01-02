/*
------------------------------------------------------------------------------------------------------------------------
####### router ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package router

import (
	"github.com/mls-361/minikit"

	"github.com/mls-361/armen/internal/components"
)

type (
	// Router AFAIRE.
	Router struct {
		*minikit.Base
		router *cRouter
	}
)

// New AFAIRE.
func New(components *components.Components) *Router {
	router := newCRouter()
	components.Router = router

	return &Router{
		Base:   minikit.NewBase("router", "router"),
		router: router,
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
