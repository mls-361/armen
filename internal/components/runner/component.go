/*
------------------------------------------------------------------------------------------------------------------------
####### runner ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package runner

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mls-361/minikit"

	"github.com/mls-361/armen/internal/client"
	"github.com/mls-361/armen/internal/components"
)

type (
	// Runner AFAIRE.
	Runner struct {
		*minikit.Base
		components *components.Components
	}
)

// New AFAIRE.
func New(components *components.Components) *Runner {
	return &Runner{
		Base:       minikit.NewBase("runner", minikit.CategoryRunner),
		components: components,
	}
}

func (cr *Runner) run(m *minikit.Manager) error {
	if err := m.BuildComponents(); err != nil {
		return err
	}

	leader := cr.components.Leader
	server := cr.components.Server

	if err := server.Start(); err != nil {
		return err
	}

	leader.Start()

	defer leader.Stop()
	defer server.Stop()

	end := make(chan os.Signal, 1)
	defer close(end)

	signal.Notify(end, os.Interrupt, syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGTERM)

	<-end

	cr.components.Logger.Info("...Stopping...") //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

	return nil
}

// Run AFAIRE.
func (cr *Runner) Run(m *minikit.Manager) error {
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
