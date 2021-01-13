/*
------------------------------------------------------------------------------------------------------------------------
####### memory ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package memory

import (
	"sync"
	"time"

	"github.com/mls-361/armen-sdk/jw"

	"github.com/mls-361/armen/internal/components"
)

type (
	lock struct {
		owner  string
		expiry time.Time
	}

	// Backend AFAIRE.
	Backend struct {
		components *components.Components
		mutex      sync.Mutex
		locks      map[string]*lock
		jobs       map[string]*jw.Job
	}
)

func New(components *components.Components) *Backend {
	return &Backend{
		components: components,
		locks:      make(map[string]*lock),
		jobs:       make(map[string]*jw.Job),
	}
}

// Build AFAIRE.
func (cb *Backend) Build() error {
	return nil
}

// Close AFAIRE.
func (cb *Backend) Close() {}

/*
######################################################################################################## @(°_°)@ #######
*/
