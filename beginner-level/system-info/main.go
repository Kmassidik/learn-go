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
}
