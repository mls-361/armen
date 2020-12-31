/*
------------------------------------------------------------------------------------------------------------------------
####### runner ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package runner

import (
	"os"

	"github.com/mls-361/component"

	"github.com/mls-361/armen/internal/client"
	"github.com/mls-361/armen/internal/components"
)

type (
	// Runner AFAIRE.
	Runner struct {
		*component.Base
		components *components.Components
	}
)

// New AFAIRE.
func New(components *components.Components) *Runner {
	return &Runner{
		Base:       component.NewBase("runner", component.CategoryRunner),
		components: components,
	}
}

// Build AFAIRE.
func (cr *Runner) Build(_ *component.Manager) error {
	cr.Built()
	return nil
}

func (cr *Runner) run(m *component.Manager) error {
	if err := m.BuildComponents(); err != nil {
		return err
	}

	// AFINIR

	return nil
}

// Run AFAIRE.
func (cr *Runner) Run(m *component.Manager) error {
	if err := m.InitializeComponents(); err != nil {
		return err
	}

	defer m.CloseComponents()

	if err := m.BuildComponent("config"); err != nil {
		return err
	}

	if len(os.Args) > 1 {
		if err := client.Execute(m, cr.components); err != nil {
			return err
		}
	}

	return cr.run(m)
}

/*
######################################################################################################## @(°_°)@ #######
*/
