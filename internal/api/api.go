/*
------------------------------------------------------------------------------------------------------------------------
####### api ####### (c) 2020-2021 mls-361 ########################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package api

import (
	"expvar"
	"net/http"

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

func (api *API) createJob(rw http.ResponseWriter, _ *http.Request) {
	rw.WriteHeader(http.StatusOK)
}

// Setup AFAIRE.
func (api *API) Setup() {
	router := api.components.CRouter

	router.Get("/debug", expvar.Handler())
	router.Get("/status", http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}))

	router.Post("/job/create", http.HandlerFunc(api.createJob))
}

// AppStopping AFAIRE.
func (api *API) AppStopping() {}

/*
######################################################################################################## @(°_°)@ #######
*/
