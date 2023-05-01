package main

import (
	"fmt"
	"page-replacement-algorithms/utils"
	"page-replacement-algorithms/utils/algorithms"
)

func main() {
	// frames
	const framesQuantity int = 30
	const pagesQuantity int = 250
	const maxPageNum int = 60
	const localityMaximumFrequency = 50
	const localityMaximumHistoryLength = 15
	const localityMaximumLength = 30
	const simulationLoops int = 1000

	// pages example in the exercise overview
	// pages := []int{1, 2, 3, 4, 1, 2, 5, 1, 2, 3, 4, 5}

	// page fault counters
	fifoFaults, optFaults, lruFaults, alruFaults, randFaults := 0, 0, 0, 0, 0

	// run simulation
	for i := 0; i < simulationLoops; i++ {
		pages := utils.GeneratePages(pagesQuantity, maxPageNum, localityMaximumFrequency, localityMaximumHistoryLength, localityMaximumLength)
		fifoFaults += algorithms.Fifo(pages, framesQuantity)
		optFaults += algorithms.Opt(pages, framesQuantity)
		lruFaults += algorithms.Lru(pages, framesQuantity)
		alruFaults += algorithms.Alru(pages, framesQuantity)
		randFaults += algorithms.Rand(pages, framesQuantity)
	}

	// calculate average page faults
	fifoFaultsAvg := float64(fifoFaults) / float64(simulationLoops)
	optFaultsAvg := float64(optFaults) / float64(simulationLoops)
	lruFaultsAvg := float64(lruFaults) / float64(simulationLoops)
	alruFaultsAvg := float64(alruFaults) / float64(simulationLoops)
	randFaultsAvg := float64(randFaults) / float64(simulationLoops)

	// print results
	fmt.Println("Average page faults:\n"+
		"FIFO: ", fifoFaultsAvg, "\n"+
		"OPT: ", optFaultsAvg, "\n"+
		"LRU: ", lruFaultsAvg, "\n"+
		"ALRU: ", alruFaultsAvg, "\n"+
		"RAND: ", randFaultsAvg)
}
