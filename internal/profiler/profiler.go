package profiler

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

type Information struct {
	Memory             uint64 `json:"memory"`
	MemoryFormatted    string `json:"memory_formatted"`
	MaxMemory          uint64 `json:"max_memory"`
	MaxMemoryFormatted string `json:"max_memory_formatted"`
	Took               int64  `json:"took"`
	TookFormatted      string `json:"took_formatted"`
}

type ToPrint struct {
	Alloc      string
	TotalAlloc string
	Sys        string
	NumGC      string
	Took       string
}

var maxAlloc uint64 = 0

// PrintUsage
// Alloc is bytes of allocated heap objects.
// "Allocated" heap objects include all reachable objects, as well as unreachable objects that the garbage collector has not yet freed.
// Specifically, Alloc increases as heap objects are allocated and decreases as the heap is swept and unreachable objects are freed.
// Sweeping occurs incrementally between GC cycles, so these two processes occur simultaneously, and as a result Alloc tends to change smoothly (in contrast with the sawtooth that is typical of stop-the-world garbage collectors).
//
// TotalAlloc is cumulative bytes allocated for heap objects.
// TotalAlloc increases as heap objects are allocated, but unlike Alloc and HeapAlloc, it does not decrease when objects are freed.
//
// Sys is the total bytes of memory obtained from the OS.
// Sys is the sum of the XSys fields below. Sys measures the virtual address space reserved by the Go runtime for the heap, stacks, and other internal data structures. It's likely that not all of the virtual address space is backed by physical memory at any given moment, though in general it all was at some point.
//
// NumGC is the number of completed GC cycles.
func PrintUsage(ctx context.Context, start time.Time) Information {
	timeElapsed := time.Since(start)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	if maxAlloc < m.Sys {
		maxAlloc = m.Sys
	}

	text := ToPrint{
		Alloc:      fmt.Sprintf("%v MiB", bToMb(m.Alloc)),
		TotalAlloc: fmt.Sprintf("%v MiB", bToMb(m.TotalAlloc)),
		Sys:        fmt.Sprintf("%v MiB", bToMb(m.Sys)),
		NumGC:      fmt.Sprintf("%v", m.NumGC),
		Took:       fmt.Sprintf("%s", timeElapsed),
	}

	fmt.Printf("%+v\n", text)

	return Information{
		Memory:             m.Sys,
		MemoryFormatted:    fmt.Sprintf("%v MiB", bToMb(m.Sys)),
		MaxMemory:          maxAlloc,
		MaxMemoryFormatted: fmt.Sprintf("%v MiB", bToMb(maxAlloc)),
		Took:               timeElapsed.Milliseconds(),
		TookFormatted:      fmt.Sprintf("%s", timeElapsed),
	}
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
