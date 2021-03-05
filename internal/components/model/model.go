/*
------------------------------------------------------------------------------------------------------------------------
####### model ####### (c) 2020-2021 mls-361 ######################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package model

import (
	"sync"
	"time"

	"github.com/mls-361/armen-sdk/message"
	"github.com/mls-361/minikit"

	"github.com/mls-361/armen/internal/components"
)

type (
	// Model AFAIRE.
	Model struct {
		*minikit.Base
		components *components.Components
		jwCh       chan<- *message.Message
		njMutex    sync.Mutex
		njTimeout  time.Time
	}
)

// New AFAIRE.
func New(components *components.Components) *Model {
	cm := &Model{
		Base:       minikit.NewBase("model", "model"),
		components: components,
	}

	components.CModel = cm

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
	cm.jwCh = cm.components.CBus.AddPublisher("jw", 10, 1)
	return nil
}

// ChannelJW AFAIRE.
func (cm *Model) ChannelJW() chan<- *message.Message {
	return cm.jwCh
}

// Clean AFAIRE.
func (cm *Model) Clean() {
	cj, cw, err := cm.components.CBackend.Clean()
	if err != nil {
		cm.components.CLogger.Error("Clean error", "reason", err.Error()) //::::::::::::::::::::::::::::::::::::::::::::
		return
	}

	cm.components.CLogger.Info("Clean", "jobs", cj, "workflows", cw) //:::::::::::::::::::::::::::::::::::::::::::::::::
}

func (cm *Model) publish(topic string, data interface{}) {
	cm.jwCh <- message.New(topic, data)
}

// Close AFAIRE.
func (cm *Model) Close() {
	close(cm.jwCh)
}

/*
######################################################################################################## @(°_°)@ #######
*/
