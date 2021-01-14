/*
------------------------------------------------------------------------------------------------------------------------
####### metrics ####### (c) 2020-2021 mls-361 ###################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package metrics

import (
	"runtime"
	"time"
)

type (
	rtMetrics struct {
		Alloc        uint64 `json:"alloc"`
		TotalAlloc   uint64 `json:"total_alloc"`
		Sys          uint64 `json:"sys"`
		Mallocs      uint64 `json:"mallocs"`
		Frees        uint64 `json:"frees"`
		LiveObjects  uint64 `json:"live_objects"`
		PauseTotal   uint64 `json:"pause_total"`
		NumGC        uint32 `json:"num_gc"`
		NumGoroutine int    `json:"num_goroutine"`
		Timestamp    int64  `json:"timestamp"`
	}
)

func (cm *Metrics) updateRuntimeMetrics() {
	var rtms runtime.MemStats
	runtime.ReadMemStats(&rtms)

	rt := cm.rtMetrics

	rt.Alloc = rtms.Alloc / 1024           // Ko
	rt.TotalAlloc = rtms.TotalAlloc / 1024 // Ko
	rt.Sys = rtms.Sys / 1024               // Ko
	rt.Mallocs = rtms.Mallocs
	rt.Frees = rtms.Frees
	rt.LiveObjects = rtms.Mallocs - rtms.Frees
	rt.PauseTotal = rtms.PauseTotalNs // En nanoseconde(s)
	rt.NumGC = rtms.NumGC
	rt.NumGoroutine = runtime.NumGoroutine()
	rt.Timestamp = time.Now().UnixNano()
}

/*
######################################################################################################## @(°_°)@ #######
*/
