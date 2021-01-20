/*
------------------------------------------------------------------------------------------------------------------------
####### components ####### (c) 2020-2021 mls-361 ################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package components

import (
	"net/http"
	"time"

	"github.com/mls-361/metrics"
	"github.com/mls-361/minikit"

	"github.com/mls-361/armen-sdk/components"
	"github.com/mls-361/armen-sdk/jw"
	"github.com/mls-361/armen-sdk/logger"
)

type (
	// Application AFAIRE.
	Application interface {
		components.Application
	}

	// Backend AFAIRE.
	Backend interface {
		AcquireLock(name, owner string, duration time.Duration) (bool, error)
		ReleaseLock(name, owner string) error
		InsertJob(job *jw.Job) (bool, error)
		NextJob() (*jw.Job, error)
		UpdateJob(job *jw.Job) error
		Workflow(id string, mustExist bool) (*jw.Workflow, error)
		InsertWorkflow(wf *jw.Workflow, job *jw.Job) error
		UpdateWorkflow(wf *jw.Workflow) error
	}

	// Bus AFAIRE.
	Bus interface {
		components.Bus
	}

	// Config AFAIRE.
	Config interface {
		components.Config
	}

	// Crypto AFAIRE.
	Crypto interface {
		components.Crypto
	}

	// Leader AFAIRE.
	Leader interface {
		components.Leader
		Start()
		Stop()
	}

	// Logger AFAIRE.
	Logger interface {
		logger.Logger
	}

	// Manager AFAIRE.
	Manager interface {
		GetComponent(category string, mustExist bool) (minikit.Component, error)
	}

	// Metrics AFAIRE.
	Metrics interface {
		metrics.Metrics
	}

	// Model AFAIRE.
	Model interface {
		components.Model
		NextJob() *jw.Job
	}

	// Router AFAIRE.
	Router interface {
		components.Router
		Handler() http.Handler
	}

	// Scheduler AFAIRE.
	Scheduler interface {
		Start()
		Stop()
	}

	// Server AFAIRE.
	Server interface {
		Port() int
		Start() error
		Stop()
	}

	// Workers AFAIRE.
	Workers interface {
		Start()
		Stop()
	}

	// Components AFAIRE.
	Components struct {
		CApplication Application
		CBackend     Backend
		CBus         Bus
		CConfig      Config
		CCrypto      Crypto
		CLeader      Leader
		CLogger      Logger
		CManager     Manager
		CMetrics     Metrics
		CModel       Model
		CRouter      Router
		CScheduler   Scheduler
		CServer      Server
		CWorkers     Workers
		unknown      []components.Component
	}
)

// New AFAIRE.
func New() *Components {
	return &Components{
		unknown: make([]components.Component, 0),
	}
}

// Application AFAIRE.
func (cs *Components) Application() components.Application {
	return cs.CApplication
}

// Bus AFAIRE.
func (cs *Components) Bus() components.Bus {
	return cs.CBus
}

// Config AFAIRE.
func (cs *Components) Config() components.Config {
	return cs.CConfig
}

// Crypto AFAIRE.
func (cs *Components) Crypto() components.Crypto {
	return cs.CCrypto
}

// Leader AFAIRE.
func (cs *Components) Leader() components.Leader {
	return cs.CLeader
}

// Logger AFAIRE.
func (cs *Components) Logger() logger.Logger {
	return cs.CLogger
}

// Model AFAIRE.
func (cs *Components) Model() components.Model {
	return cs.CModel
}

// Router AFAIRE.
func (cs *Components) Router() components.Router {
	return cs.CRouter
}

// Add AFAIRE.
func (cs *Components) Add(c components.Component) {
	cs.unknown = append(cs.unknown, c)
}

/*
######################################################################################################## @(°_°)@ #######
*/
