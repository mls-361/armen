/*
------------------------------------------------------------------------------------------------------------------------
####### main ####### (c) 2020-2021 mls-361 ######################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/mls-361/component"

	"github.com/mls-361/armen/internal/components"
	"github.com/mls-361/armen/internal/components/application"
	"github.com/mls-361/armen/internal/components/bus"
	"github.com/mls-361/armen/internal/components/config"
	"github.com/mls-361/armen/internal/components/crypto"
	"github.com/mls-361/armen/internal/components/leader"
	"github.com/mls-361/armen/internal/components/logger"
	"github.com/mls-361/armen/internal/components/model"
	"github.com/mls-361/armen/internal/components/plugins"
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
	manager := component.NewManager()

	if err := manager.AddComponents(
		app,
		bus.New(),
		config.New(cs),
		crypto.New(),
		leader.New(),
		logger.New(),
		model.New(),
		plugins.New(),
		runner.New(),
		scheduler.New(),
		server.New(),
		workers.New(),
	); err != nil {
		return app.OnError(err)
	}

	if err := manager.Run(); err != nil {
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
