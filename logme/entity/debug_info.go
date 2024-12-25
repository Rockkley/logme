package entity

type DebugInfo struct {
	NumCPU       int // the number of logical CPUs usable by the current process
	CgoCalls     int // the number of cgo calls made by the current process
	NumGoroutine int // the number of goroutines that currently exist
	Alloc        int
	TotalAlloc   int
	Sys          int
	NumGC        int
}
