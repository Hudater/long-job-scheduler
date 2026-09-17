package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var intervalSeconds int
	var jobStatus string
	workerName := "Worker"

	const (
		jobStatusPending = "StatusPending"
		jobStatusRunning = "StatusRunning"
		jobStatusDone    = "StatusDone"
	)
	// it returns the bytes
	// string, _ := fmt.Printf("Worker name: '%v' at Status: '%v' after Interval: '%v' \n", workerName, jobStatus, intervalSeconds)
	// fmt.Println(string)
	fmt.Printf("Worker '%v' ran for '%v' seconds and exited with Status '%v'\n", workerName, jobStatus, intervalSeconds)
	possibleStatusStr := fmt.Sprintf("\nPossible status includes: \n1. %v \n2. %v \n3. %v \n", jobStatusPending, jobStatusRunning, jobStatusDone)
	fmt.Println(possibleStatusStr)
	// using `fmt.Printf()` here would work but not recommended when a `%` sign appears in the string
	// possibleStatus := fmt.Sprintf("\nPossible status includes: \n1. %v \n2. %v \n3. %v \n And a random % which breaks printf", jobStatusPending, jobStatusRunning, jobStatusDone)
	// fmt.Printf(possibleStatus)

	// random int b/w 0 and 100 to proceed
	intervalSeconds = rand.IntN(100)
	// fmt.Printf("\nInterval now: %v\n", intervalSeconds)

	if intervalSeconds <= 0 {
		fmt.Println("Error: Interval less than a second")
	} else {
		possibleStatus := []string {
			jobStatusPending,
			jobStatusRunning,
			jobStatusDone,
		}
	  jobStatusIndex := rand.IntN(len(possibleStatus))
		// fmt.Println(possibleStatus[jobStatus])
		switch possibleStatus[jobStatusIndex] {
			case jobStatusPending:
				fmt.Printf("Job pending\n")
			// can also do literal matching, see
			case "StatusRunning":
				fmt.Printf("Job Running\n")
			case jobStatusDone:
				fmt.Printf("Job Done\n")
	   	default:
	        fmt.Println("Error: Job status not supported")
		}
	}
	
}
