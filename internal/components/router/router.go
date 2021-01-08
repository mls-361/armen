/*
------------------------------------------------------------------------------------------------------------------------
####### router ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package router

import (
	"expvar"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mls-361/armen-sdk/components"
	"github.com/mls-361/minikit"
)

type (
	// Router AFAIRE.
	Router struct {
		*minikit.Base
		mux *httprouter.Router
	}
)

// New AFAIRE.
func New(components *components.Components) *Router {
	mux := httprouter.New()

	mux.PanicHandler = func(rw http.ResponseWriter, _ *http.Request, _ interface{}) {
		rw.WriteHeader(http.StatusInternalServerError)
	}

	mux.Handler(http.MethodGet, "/statistics", expvar.Handler())
	mux.GET("/status", func(rw http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		rw.WriteHeader(http.StatusOK)
	})

	cr := &Router{
		Base: minikit.NewBase("router", "router"),
		mux:  mux,
	}

	components.Router = cr

	return cr
}

// ServeHTTP AFAIRE.
func (cr *Router) Handler() http.Handler {
	return cr.mux
}

// Get AFAIRE.
func (cr *Router) Get(path string, handler http.HandlerFunc) {
	cr.mux.HandlerFunc(http.MethodGet, path, handler)
}

// Post AFAIRE.
func (cr *Router) Post(path string, handler http.HandlerFunc) {
	cr.mux.HandlerFunc(http.MethodPost, path, handler)
}

/*
######################################################################################################## @(°_°)@ #######
*/
