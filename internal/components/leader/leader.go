/*
------------------------------------------------------------------------------------------------------------------------
####### leader ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package leader

import (
	"sync"
	"time"

	"github.com/mls-361/minikit"

	"github.com/mls-361/armen/internal/components"
)

const (
	_lockName    = "leader"
	_lockTimeout = 20 * time.Second
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
		Base:       minikit.NewBase("leader", ""),
		components: components,
	}

	components.CLeader = cl

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
	ok, err := cl.components.CModel.AcquireLock(
		_lockName,
		cl.components.CApplication.ID(),
		_lockTimeout,
	)
	if err != nil {
		cl.components.CLogger.Warning( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"Impossible to become the leader",
			"reason", err.Error(),
		)

		return
	}

	cl.mutex.Lock()
	defer cl.mutex.Unlock()

	if cl.amITheLeader != ok {
		cl.components.CLogger.Notice("Leader", "amITheLeader", ok) //:::::::::::::::::::::::::::::::::::::::::::::::::::
		cl.amITheLeader = ok
	}
}

// Start AFAIRE.
func (cl *Leader) Start() {
	cl.waitGroup.Add(1)

	go func() { //@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		defer cl.waitGroup.Done()

		timer := time.NewTimer(0)

		for {
			select {
			case <-timer.C:
				cl.tryAcquireLock()
			case <-cl.stopCh:
				if !timer.Stop() {
					<-timer.C
				}

				return
			}

			if cl.amITheLeader {
				timer.Reset(_lockTimeout * 3 / 4)
			} else {
				timer.Reset(_lockTimeout / 4)
			}
		}
	}()

	cl.stopCh = make(chan struct{})

	cl.components.CLogger.Info(">>>Leader") //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
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

	cl.components.CLogger.Info("<<<Leader") //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

	if cl.amITheLeader {
		if err := cl.components.CModel.ReleaseLock(_lockName, cl.components.CApplication.ID()); err != nil {
			cl.components.CLogger.Error(err.Error(), "func", "leader.Stop") //::::::::::::::::::::::::::::::::::::::::::
		}
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
