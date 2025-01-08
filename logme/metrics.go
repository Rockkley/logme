package logme

import "runtime"

type RuntimeMetrics struct {
	NumCPU       int    // the number of logical CPUs usable by the current process
	CgoCalls     int    // the number of cgo calls made by the current process
	NumGoroutine int    // the number of goroutines that currently exist
	Alloc        uint64 // bytes of allocated heap objects (in megabytes)
	TotalAlloc   uint64 // cumulative bytes allocated for heap objects (in megabytes)
	Sys          uint64 // total bytes of memory obtained from the OS (in megabytes)
	NumGC        uint32 // the number of completed GC cycles
}

func GetRuntimeMetrics() *RuntimeMetrics {
	var ms runtime.MemStats

	runtime.ReadMemStats(&ms)

	return &RuntimeMetrics{
		NumCPU:       runtime.NumCPU(),
		CgoCalls:     int(runtime.NumCgoCall()),
		NumGoroutine: runtime.NumGoroutine(),
		Alloc:        ms.Alloc / (1024 * 1024),
		TotalAlloc:   ms.TotalAlloc / (1024 * 1024),
		Sys:          ms.Sys / (1024 * 1024),
		NumGC:        ms.NumGC,
	}
}
