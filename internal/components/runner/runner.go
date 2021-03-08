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
		Base:       minikit.NewBase(minikit.CategoryRunner, ""),
		components: components,
	}
}

func (cr *Runner) waitEnd() {
	end := make(chan os.Signal, 1)

	signal.Notify(end, os.Interrupt, syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGTERM)

	<-end

	cr.components.Logger().Info("...Stopping...") //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

	close(end)
}

func (cr *Runner) run(m *minikit.Manager) error {
	if err := m.BuildComponents(); err != nil {
		return err
	}

	leader := cr.components.CLeader
	scheduler := cr.components.CScheduler
	server := cr.components.CServer
	workers := cr.components.CWorkers

	if err := server.Start(); err != nil {
		return err
	}

	workers.Start()
	leader.Start()
	scheduler.Start()

	cr.waitEnd()

	scheduler.Stop()
	leader.Stop()
	workers.Stop()
	server.Stop()

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
