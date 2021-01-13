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

	"github.com/mls-361/minikit"

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

	// Server AFAIRE.
	Server struct {
		*minikit.Base
		components *components.Components
		config     *config
		server     *http.Server
		stopped    chan error
	}
)

// New AFAIRE.
func New(components *components.Components) *Server {
	cs := &Server{
		Base:       minikit.NewBase("server", "server"),
		components: components,
		config:     &config{Port: _defaultPort},
		stopped:    make(chan error, 1),
	}

	components.CServer = cs

	return cs
}

// Dependencies AFAIRE.
func (cs *Server) Dependencies() []string {
	return []string{
		"config",
		"logger",
		"router",
	}
}

// Build AFAIRE.
func (cs *Server) Build(_ *minikit.Manager) error {
	if err := cs.components.CConfig.Decode(&cs.config, false, "components", "server"); err != nil {
		return err
	}

	cs.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", cs.config.Port),
		Handler:      cs.components.CRouter.Handler(),
		ErrorLog:     cs.components.CLogger.NewStdLogger("error", "", log.Llongfile),
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if cs.config.TLS {
		cs.server.TLSConfig = &tls.Config{
			PreferServerCipherSuites: true,
			CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
		}
	}

	return nil
}

// Port AFAIRE.
func (cs *Server) Port() int {
	return cs.config.Port
}

// Start AFAIRE.
func (cs *Server) Start() error {
	go func() { //@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		var err error

		if cs.config.TLS {
			err = cs.server.ListenAndServeTLS(cs.config.CertFile, cs.config.KeyFile)
		} else {
			err = cs.server.ListenAndServe()
		}

		cs.stopped <- err
	}()

	select {
	case err := <-cs.stopped:
		return err
	case <-time.After(50 * time.Millisecond):
		cs.components.CLogger.Info(">>>Server", "port", cs.config.Port) //::::::::::::::::::::::::::::::::::::::::::::::
		return nil
	}
}

// Stop AFAIRE.
func (cs *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cs.server.SetKeepAlivesEnabled(false)

	if err := cs.server.Shutdown(ctx); err != nil {
		cs.components.CLogger.Error(err.Error(), "func", "server.Shutdown") //::::::::::::::::::::::::::::::::::::::::::
	}

	if err := <-cs.stopped; !errors.Is(err, http.ErrServerClosed) {
		cs.components.CLogger.Error(err.Error(), "func", "server.ListenAndServe[TLS]") //:::::::::::::::::::::::::::::::
	}

	cs.components.CLogger.Info("<<<Server") //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
}

// Close AFAIRE.
func (cs *Server) Close() {
	close(cs.stopped)
}

/*
######################################################################################################## @(°_°)@ #######
*/
