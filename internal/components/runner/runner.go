/*
------------------------------------------------------------------------------------------------------------------------
####### runner ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package runner

import (
	"github.com/mls-361/component"
)

type (
	// Runner AFAIRE.
	Runner struct {
		*component.Base
	}
)

// New AFAIRE.
func New() *Runner {
	return &Runner{
		Base: component.NewBase("runner", component.CategoryRunner),
	}
}

// Build AFAIRE.
func (cr *Runner) Build(_ *component.Manager) error {
	cr.Built()
	return nil
}

// Run AFAIRE.
func (cr *Runner) Run(m *component.Manager) error {
	if err := m.InitializeComponents(); err != nil {
		return err
	}

	defer m.CloseComponents()

	if err := m.BuildComponents(); err != nil {
		return err
	}

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
