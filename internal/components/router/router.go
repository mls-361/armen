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
)

type (
	router struct {
		mux *httprouter.Router
	}
)

func newRouter() *router {
	mux := httprouter.New()

	mux.PanicHandler = func(rw http.ResponseWriter, _ *http.Request, _ interface{}) {
		rw.WriteHeader(http.StatusInternalServerError)
	}

	mux.Handler(http.MethodGet, "/statistics", expvar.Handler())
	mux.GET("/status", func(rw http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		rw.WriteHeader(http.StatusOK)
	})

	return &router{
		mux: mux,
	}
}

// ServeHTTP AFAIRE.
func (cr *router) Handler() http.Handler {
	return cr.mux
}

// Get AFAIRE.
func (cr *router) Get(path string, handler http.HandlerFunc) {
	cr.mux.HandlerFunc(http.MethodGet, path, handler)
}

// Post AFAIRE.
func (cr *router) Post(path string, handler http.HandlerFunc) {
	cr.mux.HandlerFunc(http.MethodPost, path, handler)
}

/*
######################################################################################################## @(°_°)@ #######
*/
