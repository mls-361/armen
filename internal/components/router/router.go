/*
------------------------------------------------------------------------------------------------------------------------
####### router ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type (
	router struct {
		httprouter *httprouter.Router
	}
)

func newRouter() *router {
	httprouter := httprouter.New()

	httprouter.PanicHandler = func(rw http.ResponseWriter, _ *http.Request, _ interface{}) {
		rw.WriteHeader(http.StatusInternalServerError)
	}

	return &router{
		httprouter: httprouter,
	}
}

// ServeHTTP AFAIRE.
func (cr *router) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	cr.httprouter.ServeHTTP(rw, r)
}

// Get AFAIRE.
func (cr *router) Get(path string, handler http.HandlerFunc) {
	cr.httprouter.HandlerFunc(http.MethodGet, path, handler)
}

// Post AFAIRE.
func (cr *router) Post(path string, handler http.HandlerFunc) {
	cr.httprouter.HandlerFunc(http.MethodPost, path, handler)
}

/*
######################################################################################################## @(°_°)@ #######
*/
