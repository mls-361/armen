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
	cRouter struct {
		mux *httprouter.Router
	}
)

func newCRouter() *cRouter {
	mux := httprouter.New()

	mux.PanicHandler = func(rw http.ResponseWriter, _ *http.Request, _ interface{}) {
		rw.WriteHeader(http.StatusInternalServerError)
	}

	mux.Handler(http.MethodGet, "/statistics", expvar.Handler())
	mux.GET("/status", func(rw http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		rw.WriteHeader(http.StatusOK)
	})

	return &cRouter{
		mux: mux,
	}
}

// ServeHTTP AFAIRE.
func (cr *cRouter) Handler() http.Handler {
	return cr.mux
}

// Get AFAIRE.
func (cr *cRouter) Get(path string, handler http.HandlerFunc) {
	cr.mux.HandlerFunc(http.MethodGet, path, handler)
}

// Post AFAIRE.
func (cr *cRouter) Post(path string, handler http.HandlerFunc) {
	cr.mux.HandlerFunc(http.MethodPost, path, handler)
}

/*
######################################################################################################## @(°_°)@ #######
*/
