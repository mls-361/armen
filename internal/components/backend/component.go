/*
------------------------------------------------------------------------------------------------------------------------
####### backend ####### (c) 2020-2021 mls-361 ###################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package backend

import (
	"fmt"

	"github.com/mls-361/armen-sdk/components"
	"github.com/mls-361/minikit"

	"github.com/mls-361/armen/internal/components/backend/memory"
	"github.com/mls-361/armen/internal/components/backend/pgsql"
)

type (
	backend interface {
		components.Backend
		Build() error
	}

	// Backend AFAIRE.
	Backend struct {
		*minikit.Base
		backend backend
	}
)

// New AFAIRE.
func New(components *components.Components) *Backend {
	var backend backend

	value, _ := components.Application.LookupEnv("BACKEND")

	switch value {
	case "memory":
		backend = memory.New(components)
	default:
		backend = pgsql.New(components)
	}

	if components.Application.Debug() > 1 {
		fmt.Printf("=== Backend: backend=%s\n", value) //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	}

	components.Backend = backend

	return &Backend{
		Base:    minikit.NewBase("backend", "backend"),
		backend: backend,
	}
}

// Dependencies AFAIRE.
func (cb *Backend) Dependencies() []string {
	return []string{
		"config",
		"logger",
	}
}

// Build AFAIRE.
func (cb *Backend) Build(_ *minikit.Manager) error {
	if err := cb.backend.Build(); err != nil {
		return err
	}

	cb.Built()

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
