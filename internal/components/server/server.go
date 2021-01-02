/*
------------------------------------------------------------------------------------------------------------------------
####### server ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package server

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mls-361/armen/internal/components"
)

const (
	_defaultPort = 65535
)

type (
	config struct {
		Port     int
		TLS      bool
		CertFile string
		KeyFile  string
	}

	server struct {
		components *components.Components
		config     *config
		httpserver *http.Server
		stopped    chan error
	}
)

func newServer(components *components.Components) *server {
	return &server{
		components: components,
		config:     &config{Port: _defaultPort},
		stopped:    make(chan error, 1),
	}
}

func (cs *server) build() error {
	if err := cs.components.Config.Decode(&cs.config, false, "components", "server"); err != nil {
		return err
	}

	cs.httpserver = &http.Server{
		Addr:         fmt.Sprintf(":%d", cs.config.Port),
		Handler:      cs.components.Router.Handler(),
		ErrorLog:     cs.components.Logger.NewStdLogger("error", "", log.Llongfile),
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if cs.config.TLS {
		cs.httpserver.TLSConfig = &tls.Config{
			PreferServerCipherSuites: true,
			CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
		}
	}

	return nil
}

// Port AFAIRE.
func (cs *server) Port() int {
	return cs.config.Port
}

// Start AFAIRE.
func (cs *server) Start() error {
	go func() { //@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		var err error

		if cs.config.TLS {
			err = cs.httpserver.ListenAndServeTLS(cs.config.CertFile, cs.config.KeyFile)
		} else {
			err = cs.httpserver.ListenAndServe()
		}

		cs.stopped <- err
	}()

	select {
	case err := <-cs.stopped:
		return err
	case <-time.After(50 * time.Millisecond):
		cs.components.Logger.Info(">>>Server", "port", cs.config.Port) //:::::::::::::::::::::::::::::::::::::::::::::::
		return nil
	}
}

// Stop AFAIRE.
func (cs *server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cs.httpserver.SetKeepAlivesEnabled(false)

	if err := cs.httpserver.Shutdown(ctx); err != nil {
		cs.components.Logger.Error(err.Error(), "func", "Server.Shutdown()") //:::::::::::::::::::::::::::::::::::::::::
	}

	if err := <-cs.stopped; !errors.Is(err, http.ErrServerClosed) {
		cs.components.Logger.Error(err.Error(), "func", "Server.ListenAndServe[TLS]()") //::::::::::::::::::::::::::::::
	}

	cs.components.Logger.Info("<<<Server") //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
}

func (cs *server) close() {
	close(cs.stopped)
}

/*
######################################################################################################## @(°_°)@ #######
*/
