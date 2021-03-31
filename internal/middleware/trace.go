/*
------------------------------------------------------------------------------------------------------------------------
####### middleware ####### (c) 2020-2021 mls-361 ################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package middleware

import (
	"net/http"

	"github.com/mls-361/logger"
	"github.com/mls-361/uuid"
)

type (
	responseWriter struct {
		http.ResponseWriter
		status int
	}
)

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		status:         200,
	}
}

// Trace AFAIRE.
func Trace(next http.Handler, logger logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New()
		rw := newResponseWriter(w)

		logger.Trace( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"Request",
			"id", id,
			"from", r.RemoteAddr,
			"method", r.Method,
			"uri", r.URL.RequestURI(),
		)

		next.ServeHTTP(rw, r)

		logger.Trace( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"Response",
			"id", id,
			"status", rw.status,
		)
	})
}

/*
######################################################################################################## @(°_°)@ #######
*/
