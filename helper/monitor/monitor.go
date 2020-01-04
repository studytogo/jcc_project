package monitor

import (
	"encoding/json"
	"fmt"
	"new_erp_agent_by_go/helper/error_message"

	"runtime"
	"time"
)

var startTime = time.Now()

type SysStatus struct {
	Uptime       string `json:"uptime"`
	NumGoroutine int    `json:"num_goroutine"`

	// General statistics.
	MemAllocated string `json:"mem_allocated"` // bytes allocated and still in use
	MemTotal     string `json:"mem_total"`     // bytes allocated (even if freed)
	MemSys       string `json:"mem_sys"`       // bytes obtained from system (sum of XxxSys below)
	Lookups      uint64 `json:"lookups"`       // number of pointer lookups
	MemMallocs   uint64 `json:"mem_mallocs"`   // number of mallocs
	MemFrees     uint64 `json:"mem_frees"`     // number of frees

	// Main allocation heap statistics.
	HeapAlloc    string `json:"heap_alloc"`    // bytes allocated and still in use
	HeapSys      string `json:"heap_sys"`      // bytes obtained from system
	HeapIdle     string `json:"heap_idle"`     // bytes in idle spans
	HeapInuse    string `json:"heap_inuse"`    // bytes in non-idle span
	HeapReleased string `json:"heap_released"` // bytes released to the OS
	HeapObjects  uint64 `json:"heap_objects"`  // total number of allocated objects

	// Low-level fixed-size structure allocator statistics.
	//	Inuse is bytes used now.
	//	Sys is bytes obtained from system.
	StackInuse  string `json:"stack_inuse"` // bootstrap stacks
	StackSys    string `json:"stack_sys"`
	MSpanInuse  string `json:"m_span_inuse"` // mspan structures
	MSpanSys    string `json:"m_span_sys"`
	MCacheInuse string `json:"m_cache_inuse"` // mcache structures
	MCacheSys   string `json:"m_cache_sys"`
	BuckHashSys string `json:"buck_hash_sys"` // profiling bucket hash table
	GCSys       string `json:"gc_sys"`        // GC metadata
	OtherSys    string `json:"other_sys"`     // other system allocations

	// Garbage collector statistics.
	NextGC       string `json:"next_gc"` // next run in HeapAlloc time (bytes)
	LastGC       string `json:"last_gc"` // last run in absolute time (ns)
	PauseTotalNs string `json:"pause_total_ns"`
	PauseNs      string `json:"pause_ns"` // circular buffer of recent GC pause times, most recent at [(NumGC+255)%256]
	NumGC        uint32 `json:"num_gc"`
}

func GetSystemStatus() SysStatus {
	var systemStatus SysStatus

	systemStatus.Uptime = TimeSincePro(startTime)

	m := new(runtime.MemStats)
	runtime.ReadMemStats(m)
	systemStatus.NumGoroutine = runtime.NumGoroutine()

	systemStatus.MemAllocated = fileSize(int64(m.Alloc))
	systemStatus.MemTotal = fileSize(int64(m.TotalAlloc))
	systemStatus.MemSys = fileSize(int64(m.Sys))
	systemStatus.Lookups = m.Lookups
	systemStatus.MemMallocs = m.Mallocs
	systemStatus.MemFrees = m.Frees

	systemStatus.HeapAlloc = fileSize(int64(m.HeapAlloc))
	systemStatus.HeapSys = fileSize(int64(m.HeapSys))
	systemStatus.HeapIdle = fileSize(int64(m.HeapIdle))
	systemStatus.HeapInuse = fileSize(int64(m.HeapInuse))
	systemStatus.HeapReleased = fileSize(int64(m.HeapReleased))
	systemStatus.HeapObjects = m.HeapObjects

	systemStatus.StackInuse = fileSize(int64(m.StackInuse))
	systemStatus.StackSys = fileSize(int64(m.StackSys))
	systemStatus.MSpanInuse = fileSize(int64(m.MSpanInuse))
	systemStatus.MSpanSys = fileSize(int64(m.MSpanSys))
	systemStatus.MCacheInuse = fileSize(int64(m.MCacheInuse))
	systemStatus.MCacheSys = fileSize(int64(m.MCacheSys))
	systemStatus.BuckHashSys = fileSize(int64(m.BuckHashSys))
	systemStatus.GCSys = fileSize(int64(m.GCSys))
	systemStatus.OtherSys = fileSize(int64(m.OtherSys))

	systemStatus.NextGC = fileSize(int64(m.NextGC))
	systemStatus.LastGC = fmt.Sprintf("%.1fs", float64(time.Now().UnixNano()-int64(m.LastGC))/1000/1000/1000)
	systemStatus.PauseTotalNs = fmt.Sprintf("%.1fs", float64(m.PauseTotalNs)/1000/1000/1000)
	systemStatus.PauseNs = fmt.Sprintf("%.3fs", float64(m.PauseNs[(m.NumGC+255)%256])/1000/1000/1000)
	systemStatus.NumGC = m.NumGC

	return systemStatus
}

func StandardOutput() string {
	sys := GetSystemStatus()
	if jsonStr, err := json.Marshal(sys); err != nil {
		return "MEG: " + error_message.ParMonitorErr + "INFO:" + err.Error()
	} else {
		return string(jsonStr)
	}
}
