/*
------------------------------------------------------------------------------------------------------------------------
####### model ####### (c) 2020-2021 mls-361 ######################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package supervisor

import (
	"github.com/mls-361/armen-sdk/message"
	"github.com/mls-361/minikit"

	"github.com/mls-361/armen/internal/components"
)

type (
	// Supervisor AFAIRE.
	Supervisor struct {
		*minikit.Base
		components *components.Components
	}
)

// New AFAIRE.
func New(components *components.Components) *Supervisor {
	return &Supervisor{
		Base:       minikit.NewBase("supervisor", "supervisor"),
		components: components,
	}
}

// Dependencies AFAIRE.
func (cs *Supervisor) Dependencies() []string {
	return []string{
		"bus",
		"logger",
		"model",
	}
}

// Build AFAIRE.
func (cs *Supervisor) Build(_ *minikit.Manager) error {
	return cs.components.CBus.Subscribe(
		cs.consume,
		`clean`,
	)
}

func (cs *Supervisor) consume(msg *message.Message) {
	cs.components.CLogger.Debug( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
		"Consume message",
		"id", msg.ID,
		"topic", msg.Topic,
		"publisher", msg.Publisher,
	)

	switch msg.Topic {
	case "clean":
	default:
		cs.components.CLogger.Error( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"The topic of this message is not valid",
			"id", msg.ID,
			"topic", msg.Topic,
			"publisher", msg.Publisher,
		)
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
