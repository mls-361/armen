/*
------------------------------------------------------------------------------------------------------------------------
####### application ####### (c) 2020-2021 mls-361 ################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package application

import (
	"github.com/mls-361/application"
	"github.com/mls-361/component"
)

type (
	// Application AFAIRE.
	Application struct {
		*component.Base
		application *application.Application
	}
)

// New AFAIRE.
func New(name, version, builtAt string) *Application {
	return &Application{
		Base:        component.NewBase("application", "application"),
		application: application.New(name, version, builtAt),
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

// Build AFAIRE.
func (ca *Application) Build(_ *component.Manager) error {
	ca.Built()
	ca.SetComponent(ca.application)
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
