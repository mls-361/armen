/*
------------------------------------------------------------------------------------------------------------------------
####### main ####### (c) 2020-2021 mls-361 ######################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"errors"
	"math/rand"
	"os"
	"plugin"
	"time"

	"github.com/mls-361/armen-sdk/components"
	"github.com/mls-361/failure"
	"github.com/mls-361/minikit"

	"github.com/mls-361/armen/internal/client"
	"github.com/mls-361/armen/internal/components/application"
	"github.com/mls-361/armen/internal/components/backend"
	"github.com/mls-361/armen/internal/components/bus"
	"github.com/mls-361/armen/internal/components/config"
	"github.com/mls-361/armen/internal/components/crypto"
	"github.com/mls-361/armen/internal/components/leader"
	"github.com/mls-361/armen/internal/components/logger"
	"github.com/mls-361/armen/internal/components/model"
	"github.com/mls-361/armen/internal/components/router"
	"github.com/mls-361/armen/internal/components/runner"
	"github.com/mls-361/armen/internal/components/scheduler"
	"github.com/mls-361/armen/internal/components/server"
	"github.com/mls-361/armen/internal/components/workers"
)

var (
	_version string
	_builtAt string
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func run() error {
	cs := components.New()
	app := application.New(cs, "armen", _version, _builtAt)
	manager := minikit.NewManager(app)

	if err := manager.AddComponents(
		app,
		backend.New(cs),
		bus.New(cs),
		config.New(cs),
		crypto.New(cs),
		leader.New(cs),
		logger.New(cs),
		model.New(cs),
		router.New(cs),
		runner.New(cs),
		scheduler.New(cs),
		server.New(cs),
		workers.New(),
	); err != nil {
		return app.OnError(err)
	}

	if err := manager.AddPlugins(
		app.PluginsDir(),
		"Export",
		func(_ *minikit.Manager, pSym plugin.Symbol) error {
			fn, ok := pSym.(func(*minikit.Manager, *components.Components))
			if !ok {
				return failure.New(nil).Msg("PLUGINS ERROR") // ???
			}

			fn(manager, cs)

			return nil
		},
	); err != nil {
		return app.OnError(err)
	}

	if err := manager.Run(); err != nil {
		if errors.Is(err, client.ErrStopApp) { // -help, -version & mode client ////////////////////////////////////////
			return nil
		}

		return app.OnError(err)
	}

	return nil
}

func main() {
	if run() != nil {
		os.Exit(1)
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
