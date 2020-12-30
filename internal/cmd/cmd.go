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
	// CsFs AFAIRE.
	CsFs struct {
		cs *components.Components
		fs *flag.FlagSet
	}
)

// New AFAIRE.
func New(cs *components.Components) *CsFs {
	fs := flag.NewFlagSet(cs.Application.Name(), flag.ContinueOnError)
	fs.SetOutput(os.Stdout)

	return &CsFs{
		cs: cs,
		fs: fs,
	}
}

// Components AFAIRE
func (c *CsFs) Components() *components.Components {
	return c.cs
}

// StringVar AFAIRE.
func (c *CsFs) StringVar(p *string, name string, value string, usage string) {
	c.fs.StringVar(p, name, value, usage)
}

// Usage AFAIRE.
func (c *CsFs) Usage() {
	c.fs.PrintDefaults()
}

// Parse AFAIRE.
func (c *CsFs) Parse() error {
	return c.fs.Parse(os.Args[2:])
}

/*
######################################################################################################## @(°_°)@ #######
*/
