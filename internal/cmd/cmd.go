/*
------------------------------------------------------------------------------------------------------------------------
####### cmd ####### (c) 2020-2021 mls-361 ########################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package cmd

import (
	"flag"
	"os"

	"github.com/mls-361/armen/internal/components"
)

type (
	// CmdFS AFAIRE.
	CmdFS struct {
		components *components.Components
		flagSet    *flag.FlagSet
	}
)

// New AFAIRE.
func New(components *components.Components) *CmdFS {
	flagSet := flag.NewFlagSet(components.CApplication.Name(), flag.ContinueOnError)
	flagSet.SetOutput(os.Stdout)

	return &CmdFS{
		components: components,
		flagSet:    flagSet,
	}
}

// Components AFAIRE
func (c *CmdFS) Components() *components.Components {
	return c.components
}

// StringVar AFAIRE.
func (c *CmdFS) StringVar(p *string, name string, value string, usage string) {
	c.flagSet.StringVar(p, name, value, usage)
}

// Usage AFAIRE.
func (c *CmdFS) Usage() {
	c.flagSet.PrintDefaults()
}

// Parse AFAIRE.
func (c *CmdFS) Parse() error {
	return c.flagSet.Parse(os.Args[2:])
}

/*
######################################################################################################## @(°_°)@ #######
*/
