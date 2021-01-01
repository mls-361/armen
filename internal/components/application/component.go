/*
------------------------------------------------------------------------------------------------------------------------
####### application ####### (c) 2020-2021 mls-361 ################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package application

import (
	"github.com/mls-361/application"
	"github.com/mls-361/minikit"

	"github.com/mls-361/armen/internal/components"
)

type (
	// Application AFAIRE.
	Application struct {
		*minikit.Base
		application *application.Application
	}
)

// New AFAIRE.
func New(components *components.Components, name, version, builtAt string) *Application {
	application := application.New(name, version, builtAt)
	components.Application = application

	return &Application{
		Base:        minikit.NewBase("application", "application"),
		application: application,
	}
}

// OnError AFAIRE.
func (ca *Application) OnError(err error) error {
	return ca.application.OnError(err)
}

// Initialize AFAIRE.
func (ca *Application) Initialize(_ *minikit.Manager) error {
	return ca.application.Initialize()
}

// Devel AFAIRE.
func (ca *Application) Devel() int {
	return ca.application.Devel()
}

/*
######################################################################################################## @(°_°)@ #######
*/
