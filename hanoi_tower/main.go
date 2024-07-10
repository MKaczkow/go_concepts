package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

// Recursive function to solve Tower of Hanoi
func towerOfHanoi(n int, fromRod, toRod, auxRod string, verbose bool) {
	if n == 1 {
		if verbose {
			fmt.Printf("Move disk 1 from rod %s to rod %s\n", fromRod, toRod)
		}
		return
	}
	towerOfHanoi(n-1, fromRod, auxRod, toRod, verbose)
	if verbose {
		fmt.Printf("Move disk %d from rod %s to rod %s\n", n, fromRod, toRod)
	}
	towerOfHanoi(n-1, auxRod, toRod, fromRod, verbose)
}

func main() {
	// For readability of output with large number of discs
	verbose := false

	numberOfDiscsList := []int{3, 5, 7, 10, 12, 17, 21, 32}

	// Iterate over the list of number of discs
	for _, n := range numberOfDiscsList {
		// Measure time
		start := time.Now()
		fmt.Printf("Number of disks: %d\n", n)
		towerOfHanoi(n, "A", "C", "B", verbose)
		end := time.Now()
		fmt.Printf("Time taken: %v\n", end.Sub(start))

		// Measure memory
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		fmt.Printf("Memory used: %v bytes\n", mem.TotalAlloc)

		// Write metrics to log file
		logFile, err := os.OpenFile("metrics.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer logFile.Close()

		log.SetOutput(logFile)
		log.Printf("Number of disks: %d\n", n)
		log.Printf("Time taken: %v\n", end.Sub(start))
		log.Printf("Memory used: %v bytes\n", mem.TotalAlloc)
		log.Println()
	}
}
