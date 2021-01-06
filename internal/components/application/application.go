/*
------------------------------------------------------------------------------------------------------------------------
####### application ####### (c) 2020-2021 mls-361 ################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package application

import (
	"github.com/mls-361/application"
	"github.com/mls-361/armen-sdk/components"
	"github.com/mls-361/minikit"
)

type (
	// Application AFAIRE.
	Application struct {
		*minikit.Base
		*application.Application
	}
)

// New AFAIRE.
func New(components *components.Components, name, version, builtAt string) *Application {
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

/*
######################################################################################################## @(°_°)@ #######
*/
