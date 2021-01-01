/*
------------------------------------------------------------------------------------------------------------------------
####### client ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package client

import (
	"errors"
	"fmt"
	"os"

	"github.com/mls-361/minikit"
	"github.com/mls-361/failure"

	"github.com/mls-361/armen/internal/cmd/decrypt"
	"github.com/mls-361/armen/internal/cmd/encrypt"
	"github.com/mls-361/armen/internal/components"
)

var (
	// ErrStopApp AFAIRE.
	ErrStopApp = errors.New("stop application requested")
)

type (
	cmd interface {
		Usage()
		Execute(m *minikit.Manager) error
	}
)

func initialize(cs *components.Components) map[string]cmd {
	return map[string]cmd{
		"decrypt": decrypt.New(cs),
		"encrypt": encrypt.New(cs),
	}
}

// Usage AFAIRE.
func Usage(cs *components.Components) {
	for name, cmd := range initialize(cs) {
		fmt.Println()
		fmt.Println(name)
		cmd.Usage()
	}
}

// Excecute AFAIRE.
func Execute(m *minikit.Manager, cs *components.Components) error {
	name := os.Args[1]
	c, ok := initialize(cs)[name]
	if !ok {
		return failure.New(nil).Set("command", name).Msg("this command is unknown") ////////////////////////////////////
	}

	if err := c.Execute(m); err != nil {
		return failure.New(err).
			Set("command", name).
			Msg("error when executing this command") ///////////////////////////////////////////////////////////////////
	}

	return ErrStopApp
}

/*
######################################################################################################## @(°_°)@ #######
*/
