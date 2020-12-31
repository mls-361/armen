/*
------------------------------------------------------------------------------------------------------------------------
####### server ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package server

import (
	"github.com/mls-361/component"
)

type (
	// Server AFAIRE.
	Server struct {
		*component.Base
	}
)

// New AFAIRE.
func New() *Server {
	return &Server{
		Base: component.NewBase("server", "server"),
	}
}

// Dependencies AFAIRE.
func (cs *Server) Dependencies() []string {
	return []string{
		"config",
		"logger",
	}
}

// Build AFAIRE.
func (cs *Server) Build(_ *component.Manager) error {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
