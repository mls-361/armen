/*
------------------------------------------------------------------------------------------------------------------------
####### api ####### (c) 2020-2021 mls-361 ########################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package api

import (
	"expvar"
	"net/http"

	"github.com/mls-361/armen-sdk/jw"
	"github.com/mls-361/jsonapi"

	"github.com/mls-361/armen/internal/components"
)

const (
	_maxBodySize = 1024 * 4
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

func (api *API) createJob(rw http.ResponseWriter, r *http.Request) {
	var jc *jw.JobCore

	if err := jsonapi.Decode(rw, r, _maxBodySize, &jc); err != nil {
		return
	}

	if jc.Namespace == "" {
		jsonapi.BadRequest(rw, "the namespace of the job is not specified", nil)
		return
	}

	_, err := api.components.CManager.GetComponent(jc.Namespace+".factory", true)
	if err != nil {
		jsonapi.BadRequest(rw, "the namespace of the job is not valid", err)
		return
	}

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
