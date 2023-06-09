package main

import (
	"awesome/internal/files"
	"awesome/internal/profiler"
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	now := time.Now()
	workdir := files.GetWorkdir()
	profiler.PrintUsage(ctx, now)

	file, err := os.OpenFile(workdir+"/files/custom_2017_2020.csv", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}

	profiler.PrintUsage(ctx, now)
	count, err := files.LineCounter(file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", count)
	// Output: 14897201 (lines)
	result := profiler.PrintUsage(ctx, now)
	fmt.Printf("%+v\n", result)
	// Output: {
	//  Memory: 10863824
	//  MemoryFormatted: 10
	//  MiB MaxMemory: 10863824
	//  MaxMemoryFormatted: 10
	//  MiB Took: 196
	//  TookFormatted: 196.688ms
	// }
	os.Exit(0)
}
