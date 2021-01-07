/*
------------------------------------------------------------------------------------------------------------------------
####### server ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package server

import (
	"github.com/mls-361/armen-sdk/components"
	"github.com/mls-361/minikit"
)

type (
	// Server AFAIRE.
	Server struct {
		*minikit.Base
		server *cServer
	}
)

// New AFAIRE.
func New(components *components.Components) *Server {
	server := newCServer(components)
	components.Server = server

	return &Server{
		Base:   minikit.NewBase("server", "server"),
		server: server,
	}
}

// Dependencies AFAIRE.
func (cs *Server) Dependencies() []string {
	return []string{
		"config",
		"logger",
		"router",
	}
}

// Build AFAIRE.
func (cs *Server) Build(_ *minikit.Manager) error {
	return cs.server.build()
}

// Close AFAIRE.
func (cs *Server) Close() {
	cs.server.close()
}

/*
######################################################################################################## @(°_°)@ #######
*/
