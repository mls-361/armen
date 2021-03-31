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
	"github.com/mls-361/failure"
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

	outCreateJob struct {
		ID string `json:"id"`
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
		jsonapi.BadRequest(rw, "the namespace of the job is not specified", nil) ///////////////////////////////////////
		return
	}

	c, err := api.components.CManager.GetComponent(jc.Namespace+".factory", true)
	if err != nil {
		jsonapi.BadRequest(rw, "the namespace of the job is not valid", err) ///////////////////////////////////////////
		return
	}

	factory, ok := c.(jw.Factory)
	if !ok {
		jsonapi.InternalServerError( ///////////////////////////////////////////////////////////////////////////////////
			rw,
			failure.New(nil).Set("category", c.Category()).Msg("this component is not a job factory"),
		)
	}

	if jc.Type == "" {
		jsonapi.BadRequest(rw, "the type of the job is not specified", nil) ////////////////////////////////////////////
		return
	}

	id, err := factory.CreateJob(jc)
	if err != nil {
		jsonapi.InternalServerError(rw, err) ///////////////////////////////////////////////////////////////////////////
		return
	}

	result := &outCreateJob{
		ID: id,
	}

	jsonapi.Render(rw, r, result, api.components.CLogger)
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
