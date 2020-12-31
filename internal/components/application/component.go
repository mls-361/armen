/*
------------------------------------------------------------------------------------------------------------------------
####### application ####### (c) 2020-2021 mls-361 ################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package application

import (
	"github.com/mls-361/application"
	"github.com/mls-361/component"

	"github.com/mls-361/armen/internal/components"
)

type (
	// Application AFAIRE.
	Application struct {
		*component.Base
		application *application.Application
	}
)

// New AFAIRE.
func New(components *components.Components, name, version, builtAt string) *Application {
	application := application.New(name, version, builtAt)
	components.Application = application

	return &Application{
		Base:        component.NewBase("application", "application"),
		application: application,
	}
}

// OnError AFAIRE.
func (ca *Application) OnError(err error) error {
	return ca.application.OnError(err)
}

// Initialize AFAIRE.
func (ca *Application) Initialize(_ *component.Manager) error {
	return ca.application.Initialize()
}

// Devel AFAIRE.
func (ca *Application) Devel() int {
	return ca.application.Devel()
}

// Build AFAIRE.
func (ca *Application) Build(_ *component.Manager) error {
	ca.Built()

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
