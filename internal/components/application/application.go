/*
------------------------------------------------------------------------------------------------------------------------
####### application ####### (c) 2020-2021 mls-361 ################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package application

import (
	"fmt"

	"github.com/mls-361/application"
	"github.com/mls-361/armen-sdk/components"
	"github.com/mls-361/minikit"

	"github.com/mls-361/armen/internal/util"
)

type (
	// Application AFAIRE.
	Application struct {
		*minikit.Base
		*application.Application
	}
)

// New AFAIRE.
func New(name, version, builtAt string, components *components.Components) *Application {
	application := application.New(name, version, builtAt)
	components.Application = application

	return &Application{
		Base:        minikit.NewBase("application", "application"),
		Application: application,
	}
}

// Initialize AFAIRE.
func (ca *Application) Initialize(_ *minikit.Manager) error {
	return ca.Application.Initialize()
}

// PluginsDir AFAIRE.
func (ca *Application) PluginsDir() string {
	dir, ok := ca.LookupEnv("PLUGINS")
	if !ok {
		var err error

		dir, err = util.BinaryDir()
		if err != nil {
			dir = ""
		}
	}

	if ca.Debug() > 1 {
		fmt.Printf("=== Application: pluginsDir=%s\n", dir) //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	}

	return dir
}

/*
######################################################################################################## @(°_°)@ #######
*/
