/*
------------------------------------------------------------------------------------------------------------------------
####### encrypt ####### (c) 2020-2021 mls-361 ###################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package encrypt

import (
	"fmt"

	"github.com/mls-361/component"
	"github.com/mls-361/failure"

	_cmd "github.com/mls-361/armen/internal/cmd"
	"github.com/mls-361/armen/internal/components"
)

type (
	cmd struct {
		*_cmd.CmdFS
		data struct {
			s string
		}
	}
)

// New AFAIRE.
func New(cs *components.Components) *cmd {
	return &cmd{
		CmdFS: _cmd.New(cs),
	}
}

func (c *cmd) setFlags() {
	c.StringVar(&c.data.s, "str", "", "the string to be encrypted.")
}

// Usage AFAIRE.
func (c *cmd) Usage() {
	c.setFlags()
	fmt.Println("encrypt a string.")
	c.CmdFS.Usage()
}

// Execute AFAIRE
func (c *cmd) Execute(m *component.Manager) error {
	c.setFlags()

	if err := c.Parse(); err != nil {
		return err
	}

	if c.data.s == "" {
		return failure.New(nil).
			Msg("the string to be encrypted is empty or has not been specified") ///////////////////////////////////////
	}

	if err := m.BuildComponent("crypto"); err != nil {
		return err
	}

	s, err := c.Components().Crypto.EncryptString(c.data.s)
	if err != nil {
		return err
	}

	fmt.Printf("encrypt: %s --> %s\n", c.data.s, s)

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
