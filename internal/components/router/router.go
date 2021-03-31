/*
------------------------------------------------------------------------------------------------------------------------
####### router ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mls-361/minikit"

	"github.com/mls-361/armen/internal/components"
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

	cr := &Router{
		Base: minikit.NewBase("router", ""),
		mux:  mux,
	}

	components.CRouter = cr

	return cr
}

// ServeHTTP AFAIRE.
func (cr *Router) Handler() http.Handler {
	return cr.mux
}

// Get AFAIRE.
func (cr *Router) Get(path string, handler http.Handler) {
	cr.mux.Handler(http.MethodGet, path, handler)
}

// Post AFAIRE.
func (cr *Router) Post(path string, handler http.Handler) {
	cr.mux.Handler(http.MethodPost, path, handler)
}

/*
######################################################################################################## @(°_°)@ #######
*/
