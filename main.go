package main

import (
	"fmt"
	"page-replacement-algorithms/utils"
	"page-replacement-algorithms/utils/algorithms"
	"sync"
	"time"
)

func main() {
	const framesQuantity int = 300
	const pagesQuantity int = 600
	const maxPageNum int = 600
	const localityMaximumFrequency = 100
	const localityMaximumHistoryLength = 30
	const localityMaximumLength = 20
	const simulationLoops int = 1000

	println("SequentialSimulation")
	SequentialSimulation(simulationLoops, pagesQuantity, maxPageNum, localityMaximumFrequency, localityMaximumHistoryLength, localityMaximumLength, framesQuantity)
	println()
	println("ParallelAlgorithmsSimulation")
	ParallelAlgorithmsSimulation(simulationLoops, pagesQuantity, maxPageNum, localityMaximumFrequency, localityMaximumHistoryLength, localityMaximumLength, framesQuantity)
	println()
	println("ParallelSimulationsSimulation")
	ParallelSimulationsSimulation(simulationLoops, pagesQuantity, maxPageNum, localityMaximumFrequency, localityMaximumHistoryLength, localityMaximumLength, framesQuantity)
	println()
	println("ParallelSimulationsParallelAlgorithmsGPTSimulation")
	ParallelSimulationsParallelAlgorithmsGPTSimulation(simulationLoops, pagesQuantity, maxPageNum, localityMaximumFrequency, localityMaximumHistoryLength, localityMaximumLength, framesQuantity)
	println()
	println("ParallelSimulationsParallelAlgorithmsMySimulation")
	ParallelSimulationsParallelAlgorithmsMySimulation(simulationLoops, pagesQuantity, maxPageNum, localityMaximumFrequency, localityMaximumHistoryLength, localityMaximumLength, framesQuantity)
}

func SequentialSimulation(simulationLoops int, pagesQuantity int, maxPageNum int, localityMaximumFrequency int, localityMaximumHistoryLength int, localityMaximumLength int, framesQuantity int) {
	fifoFaults, optFaults, lruFaults, alruFaults, randFaults := 0, 0, 0, 0, 0
	start := time.Now()
	for i := 0; i < simulationLoops; i++ {
		pages := utils.GeneratePages(pagesQuantity, maxPageNum, localityMaximumFrequency, localityMaximumHistoryLength, localityMaximumLength)
		fifoFaults += algorithms.Fifo(pages, framesQuantity)
		optFaults += algorithms.Opt(pages, framesQuantity)
		lruFaults += algorithms.Lru(pages, framesQuantity)
		alruFaults += algorithms.Alru(pages, framesQuantity)
		randFaults += algorithms.Rand(pages, framesQuantity)
	}
	PrintStatistics(fifoFaults, simulationLoops, optFaults, lruFaults, alruFaults, randFaults, start)
}

func ParallelAlgorithmsSimulation(simulationLoops int, pagesQuantity int, maxPageNum int, localityMaximumFrequency int, localityMaximumHistoryLength int, localityMaximumLength int, framesQuantity int) {
	fifoFaults, optFaults, lruFaults, alruFaults, randFaults := 0, 0, 0, 0, 0
	start := time.Now()

	// channels to signal completion of each simulation loop
	fifoDone := make(chan bool)
	optDone := make(chan bool)
	lruDone := make(chan bool)
	alruDone := make(chan bool)
	randDone := make(chan bool)

	// run simulation
	for i := 0; i < simulationLoops; i++ {
		pages := utils.GeneratePages(pagesQuantity, maxPageNum, localityMaximumFrequency, localityMaximumHistoryLength, localityMaximumLength)

		// run each algorithm in a goroutine
		go func() {
			fifoFaults += algorithms.Fifo(pages, framesQuantity)
			fifoDone <- true
		}()
		go func() {
			optFaults += algorithms.Opt(pages, framesQuantity)
			optDone <- true
		}()
		go func() {
			lruFaults += algorithms.Lru(pages, framesQuantity)
			lruDone <- true
		}()
		go func() {
			alruFaults += algorithms.Alru(pages, framesQuantity)
			alruDone <- true
		}()
		go func() {
			randFaults += algorithms.Rand(pages, framesQuantity)
			randDone <- true
		}()
	}

	// wait for each algorithm to finish before calculating average page faults
	for i := 0; i < simulationLoops; i++ {
		<-fifoDone
		<-optDone
		<-lruDone
		<-alruDone
		<-randDone
	}

	PrintStatistics(fifoFaults, simulationLoops, optFaults, lruFaults, alruFaults, randFaults, start)
}

func ParallelSimulationsSimulation(simulationLoops int, pagesQuantity int, maxPageNum int, localityMaximumFrequency int, localityMaximumHistoryLength int, localityMaximumLength int, framesQuantity int) {
	fifoFaults, optFaults, lruFaults, alruFaults, randFaults := 0, 0, 0, 0, 0
	start := time.Now()

	wg := sync.WaitGroup{}
	// run simulation
	for i := 0; i < simulationLoops; i++ {
		// increment the wait group counter for each simulation loop
		wg.Add(1)

		// run each algorithm in a goroutine
		go func() {
			pages := utils.GeneratePages(pagesQuantity, maxPageNum, localityMaximumFrequency, localityMaximumHistoryLength, localityMaximumLength)

			// run each algorithm and add the faults to their respective variables
			fifoFaults += algorithms.Fifo(pages, framesQuantity)
			optFaults += algorithms.Opt(pages, framesQuantity)
			lruFaults += algorithms.Lru(pages, framesQuantity)
			alruFaults += algorithms.Alru(pages, framesQuantity)
			randFaults += algorithms.Rand(pages, framesQuantity)

			// signal completion of the simulation loop
			wg.Done()
		}()
	}

	// wait for all simulation loops to finish before calculating average page faults
	wg.Wait()

	PrintStatistics(fifoFaults, simulationLoops, optFaults, lruFaults, alruFaults, randFaults, start)
}

func ParallelSimulationsParallelAlgorithmsGPTSimulation(simulationLoops int, pagesQuantity int, maxPageNum int, localityMaximumFrequency int, localityMaximumHistoryLength int, localityMaximumLength int, framesQuantity int) {
	fifoFaults, optFaults, lruFaults, alruFaults, randFaults := 0, 0, 0, 0, 0
	start := time.Now()

	// channels to signal completion of each simulation loop and each algorithm
	simDone := make(chan bool, simulationLoops)
	fifoDone := make(chan int)
	optDone := make(chan int)
	lruDone := make(chan int)
	alruDone := make(chan int)
	randDone := make(chan int)

	// run each simulation loop in a goroutine
	for i := 0; i < simulationLoops; i++ {
		go func() {
			pages := utils.GeneratePages(pagesQuantity, maxPageNum, localityMaximumFrequency, localityMaximumHistoryLength, localityMaximumLength)

			// run each algorithm in a goroutine
			go func() {
				fifoFaults += algorithms.Fifo(pages, framesQuantity)
				fifoDone <- 1
			}()
			go func() {
				optFaults += algorithms.Opt(pages, framesQuantity)
				optDone <- 1
			}()
			go func() {
				lruFaults += algorithms.Lru(pages, framesQuantity)
				lruDone <- 1
			}()
			go func() {
				alruFaults += algorithms.Alru(pages, framesQuantity)
				alruDone <- 1
			}()
			go func() {
				randFaults += algorithms.Rand(pages, framesQuantity)
				randDone <- 1
			}()

			// wait for all algorithms to finish before signalling completion of simulation loop
			for j := 0; j < 5; j++ {
				select {
				case <-fifoDone:
				case <-optDone:
				case <-lruDone:
				case <-alruDone:
				case <-randDone:
				}
			}
			simDone <- true
		}()
	}

	// wait for all simulation loops to finish before calculating average page faults
	for i := 0; i < simulationLoops; i++ {
		<-simDone
	}

	PrintStatistics(fifoFaults, simulationLoops, optFaults, lruFaults, alruFaults, randFaults, start)
}

func ParallelSimulationsParallelAlgorithmsMySimulation(simulationLoops int, pagesQuantity int, maxPageNum int, localityMaximumFrequency int, localityMaximumHistoryLength int, localityMaximumLength int, framesQuantity int) {
	fifoFaults, optFaults, lruFaults, alruFaults, randFaults := 0, 0, 0, 0, 0
	start := time.Now()

	wg := sync.WaitGroup{}
	wg.Add(5 * simulationLoops)
	// run simulation
	for i := 0; i < simulationLoops; i++ {

		// run each algorithm in a goroutine
		go func() {
			pages := utils.GeneratePages(pagesQuantity, maxPageNum, localityMaximumFrequency, localityMaximumHistoryLength, localityMaximumLength)

			// run each algorithm and add the faults to their respective variables
			go func() {
				fifoFaults += algorithms.Fifo(pages, framesQuantity)
				wg.Done()
			}()
			go func() {
				optFaults += algorithms.Opt(pages, framesQuantity)
				wg.Done()
			}()
			go func() {
				lruFaults += algorithms.Lru(pages, framesQuantity)
				wg.Done()
			}()
			go func() {
				alruFaults += algorithms.Alru(pages, framesQuantity)
				wg.Done()
			}()
			go func() {
				randFaults += algorithms.Rand(pages, framesQuantity)
				wg.Done()
			}()
		}()
	}

	// wait for all simulation loops to finish before calculating average page faults
	wg.Wait()

	PrintStatistics(fifoFaults, simulationLoops, optFaults, lruFaults, alruFaults, randFaults, start)
}

func PrintStatistics(fifoFaults int, simulationLoops int, optFaults int, lruFaults int, alruFaults int, randFaults int, start time.Time) {
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
	fmt.Printf("Time took: %.2fs\n", time.Now().Sub(start).Seconds())
}
