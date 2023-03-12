package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
)

const (
	maxNumberOfBranchs  = 3
	maxBranchLevel = 2
)

var (
	counterBranches   atomic.Uint32
	currentBranchLevel int
	strParentNumber string
)

func main() {
	var wg sync.WaitGroup

	for BranchNumber := 1; BranchNumber <= maxNumberOfBranchs; BranchNumber++ {
		wg.Add(1)
		go goBranch(strParentNumber, BranchNumber, maxBranchLevel, currentBranchLevel, maxNumberOfBranchs, &wg)
	}

	wg.Wait()

	fmt.Printf("Total branches numbers: %v\n", counterBranches.Load())
}

func goBranch(strParentNumber string, BranchNumber, maxBranchLevel, currentBranchLevel, numberOfBranchs int, wg *sync.WaitGroup) {
	counterBranches.Add(1)	
	newBranchLevel := currentBranchLevel + 1
	
	if strParentNumber != "" {
		strParentNumber = strParentNumber + "." + strconv.Itoa(BranchNumber)
	} else {
		strParentNumber =  strconv.Itoa(BranchNumber)
	}

	fmt.Printf("Level: %v,\tBranchNumber: %v\n", newBranchLevel, strParentNumber)
	
	if newBranchLevel < maxBranchLevel {
		for i := 1; i <= numberOfBranchs; i++ {
			wg.Add(1)
			go goBranch(strParentNumber, i, maxBranchLevel, newBranchLevel, numberOfBranchs, wg)
		}
	}
	wg.Done()
}
