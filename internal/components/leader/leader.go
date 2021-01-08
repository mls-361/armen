/*
------------------------------------------------------------------------------------------------------------------------
####### leader ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package leader

import (
	"sync"
	"time"

	"github.com/mls-361/armen-sdk/components"
	"github.com/mls-361/minikit"
)

const (
	lockName    = "leader"
	lockTimeout = 20 * time.Second
)

type (
	// Leader AFAIRE.
	Leader struct {
		*minikit.Base
		components   *components.Components
		mutex        sync.Mutex
		amITheLeader bool
		waitGroup    sync.WaitGroup
		stopCh       chan struct{}
	}
)

// New AFAIRE.
func New(components *components.Components) *Leader {
	cl := &Leader{
		Base:       minikit.NewBase("leader", "leader"),
		components: components,
	}

	components.Leader = cl

	return cl
}

// Dependencies AFAIRE.
func (cl *Leader) Dependencies() []string {
	return []string{
		"application",
		"logger",
		"model",
	}
}

func (cl *Leader) tryAcquireLock() {
	ok, err := cl.components.Model.AcquireLock(
		lockName,
		cl.components.Application.ID(),
		lockTimeout,
	)
	if err != nil {
		cl.components.Logger.Warning( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"Impossible to become the leader",
			"reason", err.Error(),
		)

		return
	}

	cl.mutex.Lock()
	defer cl.mutex.Unlock()

	if cl.amITheLeader != ok {
		cl.components.Logger.Notice("Leader", "amITheLeader", ok) //::::::::::::::::::::::::::::::::::::::::::::::::::::
		cl.amITheLeader = ok
	}
}

// Start AFAIRE.
func (cl *Leader) Start() {
	cl.waitGroup.Add(1)

	go func() { //@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		defer cl.waitGroup.Done()

		var after time.Duration

		cl.tryAcquireLock()

		for {
			if cl.amITheLeader {
				after = lockTimeout * 3 / 4
			} else {
				after = lockTimeout / 4
			}

			select {
			case <-cl.stopCh:
				return
			case <-time.After(after):
				cl.tryAcquireLock()
			}
		}
	}()

	cl.stopCh = make(chan struct{})

	cl.components.Logger.Info(">>>Leader") //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
}

// AmITheLeader AFAIRE.
func (cl *Leader) AmITheLeader() bool {
	cl.mutex.Lock()
	defer cl.mutex.Unlock()

	return cl.amITheLeader
}

// Stop AFAIRE.
func (cl *Leader) Stop() {
	close(cl.stopCh)
	cl.waitGroup.Wait()

	cl.components.Logger.Info("<<<Leader") //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

	if cl.amITheLeader {
		if err := cl.components.Model.ReleaseLock(lockName, cl.components.Application.ID()); err != nil {
			cl.components.Logger.Error(err.Error(), "func", "leader.Stop") //:::::::::::::::::::::::::::::::::::::::::::
		}
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
