/*
------------------------------------------------------------------------------------------------------------------------
####### runner ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package runner

import (
	"github.com/mls-361/component"
	"github.com/mls-361/failure"
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

func (cr *Runner) initializeComponents(m *component.Manager) error {
	for _, c := range m.Components() {
		if c.IsBuilt() {
			continue
		}

		if err := c.Initialize(m); err != nil {
			return failure.New(err).Set("component", c.Name()).Msg("initialization error") /////////////////////////////
		}
	}

	return nil
}

func (cr *Runner) closeComponents(_ *component.Manager) {
}

func (cr *Runner) recursiveBuild(m *component.Manager, snitch map[string]bool, c component.Component) error {
	//fmt.Println(c.Name(), "TODO")
	snitch[c.Category()] = false

	for _, cc := range c.Dependencies() {
		d, err := m.GetComponent(cc, true)
		if err != nil {
			return err
		}

		if d.IsBuilt() {
			continue
		}

		done, ok := snitch[cc]
		if ok {
			if done {
				continue
			}

			return failure.New(nil).
				Set("component1", c.Name()).
				Set("component2", d.Name()).
				Msg("these two components are interdependent") /////////////////////////////////////////////////////////
		}

		if err := cr.recursiveBuild(m, snitch, d); err != nil {
			return err
		}
	}

	if err := c.Build(m); err != nil {
		return failure.New(err).Set("component", c.Name()).Msg("build error") //////////////////////////////////////////
	}

	//fmt.Println(c.Name(), "DONE")
	snitch[c.Category()] = true

	return nil
}

func (cr *Runner) buildComponents(m *component.Manager) error {
	snitch := make(map[string]bool)

	for _, c := range m.Components() {
		if c.IsBuilt() {
			continue
		}

		done, ok := snitch[c.Category()]
		if ok {
			if !done {
				return failure.New(nil).
					Set("component", c.Name()).
					Msg("this error should not occur") /////////////////////////////////////////////////////////////////
			}

			continue
		}

		if err := cr.recursiveBuild(m, snitch, c); err != nil {
			return err
		}
	}

	return nil
}

// Run AFAIRE.
func (cr *Runner) Run(m *component.Manager) error {
	if err := cr.initializeComponents(m); err != nil {
		return err
	}

	defer cr.closeComponents(m)

	if err := cr.buildComponents(m); err != nil {
		return err
	}

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
