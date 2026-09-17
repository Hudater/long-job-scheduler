package main

import (
	"fmt"
	"math/rand/v2"
)

const (
	jobStatusPending = "StatusPending"
	jobStatusRunning = "StatusRunning"
	jobStatusDone    = "StatusDone"
)

func main() {
	var intervalSeconds int
	// var jobStatus string
	workerName := "WorkerOne"

	possibleStatus := []string{
		jobStatusPending,
		jobStatusRunning,
		jobStatusDone,
	}
	
	// random values to proceed
	intervalSeconds = rand.IntN(100)
	// fmt.Printf("\nInterval now: %v\n", intervalSeconds)
	jobStatusIndex := rand.IntN(len(possibleStatus))
	// fmt.Printf("Job Status: %v", jobStatus)

	jobStatusStr, jobStatusBool := DescribeJobStatus(possibleStatus[jobStatusIndex])
	if jobStatusBool {
		fmt.Println(jobStatusStr)
	} else {
		fmt.Println("Error: Job status could not be found. Workername, JobStatus must be non-empty strings")
	}

	workerDurationStr, workerDurationBool := GetWorkerDuration(workerName, intervalSeconds)
	if workerDurationBool {
		fmt.Println(workerDurationStr)
	} else {
		fmt.Println("Error: Worker duration could not be found. Workername must be non-empty string and Interval for worker must be non-zero positive integer")
	}
}

func DescribeJobStatus(jobStatus string) (string, bool) {
	if jobStatus == "" {
		return "", false
	}
	switch jobStatus {
	case jobStatusPending:
		return "Job Pending", true
	case jobStatusRunning:
		return "Job Running", true
	case jobStatusDone:
		return "Job Done", true
	default:
		return "", false
	}
}

func GetWorkerDuration(workerName string, intervalSeconds int) (string, bool) {
	if intervalSeconds <= 0 || workerName == "" {
		return "", false
	}
	workerDurationFormattedString := fmt.Sprintf("Worker '%v' ran for '%v' seconds", workerName, intervalSeconds)
	return workerDurationFormattedString, true
}
