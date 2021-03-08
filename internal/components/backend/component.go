/*
------------------------------------------------------------------------------------------------------------------------
####### backend ####### (c) 2020-2021 mls-361 ###################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package backend

import (
	"fmt"

	"github.com/mls-361/minikit"

	"github.com/mls-361/armen/internal/components"
	"github.com/mls-361/armen/internal/components/backend/memory"
	"github.com/mls-361/armen/internal/components/backend/pgsql"
)

type (
	backend interface {
		components.Backend
		Build() error
		Close()
	}

	// Backend AFAIRE.
	Backend struct {
		*minikit.Base
		backend backend
	}
)

// New AFAIRE.
func New(components *components.Components) *Backend {
	var cb backend

	value, _ := components.CApplication.LookupEnv("BACKEND")

	switch value {
	case "memory":
		cb = memory.New(components)
	default:
		cb = pgsql.New(components)
	}

	if components.CApplication.Debug() > 1 {
		fmt.Printf("=== Backend: backend=%s\n", value) //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	}

	components.CBackend = cb

	return &Backend{
		Base:    minikit.NewBase("backend", ""),
		backend: cb,
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
	return cb.backend.Build()
}

// Close AFAIRE.
func (cb *Backend) Close() {
	cb.backend.Close()
}

/*
######################################################################################################## @(°_°)@ #######
*/
