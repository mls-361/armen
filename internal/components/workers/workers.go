/*
------------------------------------------------------------------------------------------------------------------------
####### workers ####### (c) 2020-2021 mls-361 ###################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package workers

import (
	"sync"

	"github.com/mls-361/metrics"
	"github.com/mls-361/minikit"

	"github.com/mls-361/armen/internal/components"
	"github.com/mls-361/armen/internal/components/workers/worker"
)

const (
	_defaultPoolSize = 10
	_maxPoolSize     = 100
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
		components  *components.Components
		config      *config
		mutex       sync.Mutex
		workers     []chan struct{}
		waitGroup   sync.WaitGroup
		mcsPoolSize metrics.GaugeInt
	}
)

// New AFAIRE.
func New(components *components.Components) *Workers {
	cw := &Workers{
		Base:       minikit.NewBase("workers", "workers"),
		components: components,
		config:     &config{Pool: struct{ Size int }{Size: _defaultPoolSize}},
		workers:    make([]chan struct{}, 0),
	}

	components.CWorkers = cw

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
	if err := cw.components.CConfig.Decode(&cw.config, false, "components", "workers"); err != nil {
		return err
	}

	if cw.config.Pool.Size < 0 {
		cw.config.Pool.Size = 0
	} else if cw.config.Pool.Size > _maxPoolSize {
		cw.config.Pool.Size = _maxPoolSize
	}

	cw.mcsPoolSize = cw.components.CMetrics.NewGaugeInt("workers.pool.size")

	return nil
}

func (cw *Workers) startWorker() {
	cw.waitGroup.Add(1)

	stopCh := make(chan struct{})
	cw.workers = append(cw.workers, stopCh)

	go func() { //@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		worker.New(cw.components, stopCh).Run()
		cw.waitGroup.Done()
	}()
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

	for len(cw.workers) < size {
		cw.startWorker()
	}

	for len(cw.workers) > size {
		cw.stopWorker()
	}

	cw.mcsPoolSize.Set(int64(size))
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

/*
######################################################################################################## @(°_°)@ #######
*/
