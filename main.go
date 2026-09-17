package main

import (
	"fmt"
)

func main() {
	var intervalSeconds int
	var jobStatus string
	workerName := "Worker"

	const (
		jobStatusPending = "StatusPending"
		jobStatusRunning = "StatusRunning"
		jobStatusDone = "StatusDone"
	)
 // it returns the bytes
	// string, _ := fmt.Printf("Worker name: '%v' at Status: '%v' after Interval: '%v' \n", workerName, jobStatus, intervalSeconds)
	// fmt.Println(string)
	fmt.Printf("Worker '%v' ran for '%v' seconds and exited with Status '%v'\n", workerName, jobStatus, intervalSeconds)
	possibleStatus := fmt.Sprintf("\nPossible status includes: \n1. %v \n2. %v \n3. %v \n", jobStatusPending, jobStatusRunning, jobStatusDone)
	fmt.Println(possibleStatus)
	// using `fmt.Printf()` here would work but not recommended when a `%` sign appears in the string
	// possibleStatus := fmt.Sprintf("\nPossible status includes: \n1. %v \n2. %v \n3. %v \n And a random % which breaks printf", jobStatusPending, jobStatusRunning, jobStatusDone)
	// fmt.Printf(possibleStatus)
	
}
