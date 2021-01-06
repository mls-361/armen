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
)

const (
	_lockName    = "leader"
	_lockTimeout = 20 * time.Second
)

type (
	cLeader struct {
		components   *components.Components
		mutex        sync.Mutex
		amITheLeader bool
		waitGroup    sync.WaitGroup
		stop         chan struct{}
	}
)

func newCLeader(components *components.Components) *cLeader {
	return &cLeader{
		components: components,
		stop:       make(chan struct{}),
	}
}

func (cl *cLeader) tryAcquireLock() {
	ok, err := cl.components.Model.AcquireLock(
		_lockName,
		cl.components.Application.ID(),
		_lockTimeout,
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
func (cl *cLeader) Start() {
	cl.waitGroup.Add(1)

	go func() { //@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		defer cl.waitGroup.Done()

		var after time.Duration

		cl.tryAcquireLock()

		for {
			if cl.amITheLeader {
				after = _lockTimeout * 3 / 4
			} else {
				after = _lockTimeout / 4
			}

			select {
			case <-cl.stop:
				return
			case <-time.After(after):
				cl.tryAcquireLock()
			}
		}
	}()

	cl.components.Logger.Info(">>>Leader") //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
}

// AmITheLeader AFAIRE.
func (cl *cLeader) AmITheLeader() bool {
	cl.mutex.Lock()
	defer cl.mutex.Unlock()

	return cl.amITheLeader
}

// Stop AFAIRE.
func (cl *cLeader) Stop() {
	close(cl.stop)
	cl.waitGroup.Wait()

	cl.components.Logger.Info("<<<Leader") //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

	if cl.amITheLeader {
		if err := cl.components.Model.ReleaseLock(_lockName, cl.components.Application.ID()); err != nil {
			cl.components.Logger.Error(err.Error(), "func", "leader.Stop") //:::::::::::::::::::::::::::::::::::::::::::
		}
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
