package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("======== System Information Report ========")

	// Fetch OS and Architecture Information
	if os, arch := runtime.GOOS, runtime.GOARCH; os != "" && arch != "" {
		fmt.Println("Operating System:", os)
		fmt.Println("Architecture:", arch)
	} else {
		fmt.Println("Error: Unable to fetch OS or architecture information")
	}

	// Fetch CPU Core Information
	numCPU := runtime.NumCPU()
	if numCPU > 0 {
		fmt.Println("CPU Cores:", numCPU)
	} else {
		fmt.Println("Error: Unable to fetch CPU core information")
	}

	// Fetch initial memory usage
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Initial Memory Allocation: %v MB\n", memStats.Alloc/1024/1024)

	// Simulate memory usage by allocating a large slice
	sampleSlice := make([]int, 1_000_000)
	sampleSlice[0] = 42 // Prevent Go to remove this cause this compiler still use

	// Fetch memory usage after simulation
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Memory Allocation After Simulation: %v MB\n", memStats.Alloc/1024/1024)
}

// you can more the memory docs in here : https://pkg.go.dev/runtime#ReadMemStats
