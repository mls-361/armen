/*
------------------------------------------------------------------------------------------------------------------------
####### model ####### (c) 2020-2021 mls-361 ######################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package model

import (
	"github.com/mls-361/armen-sdk/components"
	"github.com/mls-361/armen-sdk/message"
	"github.com/mls-361/minikit"
)

type (
	// Model AFAIRE.
	Model struct {
		*minikit.Base
		components *components.Components
		busCh      chan<- *message.Message
	}
)

// New AFAIRE.
func New(components *components.Components) *Model {
	cm := &Model{
		Base:       minikit.NewBase("model", "model"),
		components: components,
	}

	components.Model = cm

	return cm
}

// Dependencies AFAIRE.
func (cm *Model) Dependencies() []string {
	return []string{
		"backend",
		"bus",
		"logger",
	}
}

// Build AFAIRE.
func (cm *Model) Build(_ *minikit.Manager) error {
	cm.busCh = cm.components.Bus.AddPublisher("model", 1, 1)

	return nil
}

func (cm *Model) publish(topic string, data interface{}) {
	cm.busCh <- message.New(topic, data)
}

// Close AFAIRE.
func (cm *Model) Close() {
	close(cm.busCh)
}

/*
######################################################################################################## @(°_°)@ #######
*/
