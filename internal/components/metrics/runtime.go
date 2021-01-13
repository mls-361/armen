/*
------------------------------------------------------------------------------------------------------------------------
####### metrics ####### (c) 2020-2021 mls-361 ###################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package metrics

import "runtime"

type (
	rtMetrics struct {
		NumGoroutine int `json:"num_goroutine"`
	}
)

func (cm *Metrics) updateRuntimeMetrics() {
	cm.rtMetrics.NumGoroutine = runtime.NumGoroutine()
}

/*
######################################################################################################## @(°_°)@ #######
*/