/*
------------------------------------------------------------------------------------------------------------------------
####### server ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package server

import (
	"github.com/mls-361/minikit"
)

type (
	// Server AFAIRE.
	Server struct {
		*minikit.Base
	}
)

// New AFAIRE.
func New() *Server {
	return &Server{
		Base: minikit.NewBase("server", "server"),
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
func (cs *Server) Build(_ *minikit.Manager) error {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
