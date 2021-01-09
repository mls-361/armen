/*
------------------------------------------------------------------------------------------------------------------------
####### workers ####### (c) 2020-2021 mls-361 ###################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package workers

import (
	"expvar"
	"sync"

	"github.com/mls-361/armen-sdk/components"
	"github.com/mls-361/armen-sdk/message"
	"github.com/mls-361/minikit"

	"github.com/mls-361/armen/internal/components/workers/worker"
)

const (
	defaultPoolSize = 10
	maxPoolSize     = 100
)

type (
	config struct {
		Pool struct {
			Size int
		}
	}

	// Workers AFAIRE.
	Workers struct {
		*minikit.Base
		components *components.Components
		config     *config
		mutex      sync.Mutex
		workers    []chan struct{}
		busCh      chan<- *message.Message
		waitGroup  sync.WaitGroup
		stats      *expvar.Map
	}
)

// New AFAIRE.
func New(components *components.Components) *Workers {
	cw := &Workers{
		Base:       minikit.NewBase("workers", "workers"),
		components: components,
		config:     &config{Pool: struct{ Size int }{Size: defaultPoolSize}},
		workers:    make([]chan struct{}, 0),
		stats:      expvar.NewMap("workers"),
	}

	components.Workers = cw

	return cw
}

// Dependencies AFAIRE.
func (cw *Workers) Dependencies() []string {
	return []string{
		"bus",
		"config",
		"logger",
		"model",
	}
}

// Build AFAIRE.
func (cw *Workers) Build(_ *minikit.Manager) error {
	if err := cw.components.Config.Decode(&cw.config, false, "components", "workers"); err != nil {
		return err
	}

	if cw.config.Pool.Size < 0 {
		cw.config.Pool.Size = 0
	} else if cw.config.Pool.Size > maxPoolSize {
		cw.config.Pool.Size = maxPoolSize
	}

	cw.busCh = cw.components.Bus.AddPublisher("workers", 1, 1)

	return nil
}

func (cw *Workers) worker(stopCh chan struct{}) {
	worker := worker.New(
		cw.components,
		cw.busCh,
		stopCh,
		func(id string) {
			cw.stats.Add("Jobs", 1)
		},
	)

	worker.Run()

	cw.waitGroup.Done()
}

func (cw *Workers) startWorker() {
	cw.waitGroup.Add(1)

	stopCh := make(chan struct{})
	cw.workers = append(cw.workers, stopCh)

	go cw.worker(stopCh) //@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
}

func (cw *Workers) stopWorker() {
	stopCh := cw.workers[0]
	cw.workers = cw.workers[1:]
	close(stopCh)
}

// Resize AFAIRE.
func (cw *Workers) Resize(size int) {
	cw.mutex.Lock()
	defer cw.mutex.Unlock()

	delta := int64(size - len(cw.workers))

	for len(cw.workers) < size {
		cw.startWorker()
	}

	for len(cw.workers) > size {
		cw.stopWorker()
	}

	cw.stats.Add("PoolSize", delta)
}

// Start AFAIRE.
func (cw *Workers) Start() {
	cw.Resize(cw.config.Pool.Size)
}

// Stop AFAIRE.
func (cw *Workers) Stop() {
	cw.Resize(0)
	cw.waitGroup.Wait()
}

// Close AFAIRE.
func (cw *Workers) Close() {
	close(cw.busCh)
}

/*
######################################################################################################## @(°_°)@ #######
*/
