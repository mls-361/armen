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

	"github.com/mls-361/logger"
	"github.com/mls-361/minikit"
	"github.com/mls-361/uuid"

	"github.com/mls-361/armen/internal/components"
	"github.com/mls-361/armen/internal/middleware"
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
		logger     logger.Logger
		config     *config
		server     *http.Server
		errCh      chan error
	}
)

// New AFAIRE.
func New(components *components.Components) *Server {
	cs := &Server{
		Base:       minikit.NewBase("server", ""),
		components: components,
		config:     &config{Port: _defaultPort},
		errCh:      make(chan error, 1),
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

func (cs *Server) handler() http.Handler {
	return middleware.Trace(cs.components.CRouter.Handler(), cs.logger)
}

// Build AFAIRE.
func (cs *Server) Build(_ *minikit.Manager) error {
	logger := cs.components.CLogger.CreateLogger(uuid.New(), "server")

	cs.logger = logger

	if err := cs.components.CConfig.Decode(&cs.config, false, "components", "server"); err != nil {
		return err
	}

	cs.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", cs.config.Port),
		Handler:      cs.handler(),
		ErrorLog:     logger.NewStdLogger("error", "", log.Llongfile),
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

		cs.errCh <- err
	}()

	timer := time.NewTimer(50 * time.Millisecond)
	defer timer.Stop()

	select {
	case <-timer.C:
		cs.logger.Info(">>>Server", "port", cs.config.Port) //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
		return nil
	case err := <-cs.errCh:
		return err
	}
}

// Stop AFAIRE.
func (cs *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cs.server.SetKeepAlivesEnabled(false)

	if err := cs.server.Shutdown(ctx); err != nil {
		cs.logger.Error(err.Error(), "func", "server.Shutdown") //::::::::::::::::::::::::::::::::::::::::::::::::::::::
	}

	if err := <-cs.errCh; !errors.Is(err, http.ErrServerClosed) {
		cs.logger.Error(err.Error(), "func", "server.ListenAndServe[TLS]") //:::::::::::::::::::::::::::::::::::::::::::
	}

	cs.logger.Info("<<<Server") //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
}

// Close AFAIRE.
func (cs *Server) Close() {
	close(cs.errCh)
}

/*
######################################################################################################## @(°_°)@ #######
*/
