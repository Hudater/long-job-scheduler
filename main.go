package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"slices"
	"time"
)

type Job struct {
	Name        string
	Interval    int
	MaxDuration int
	Status      string
	Task        Task
}

const (
	jobStatusPending  = "StatusPending"
	jobStatusRunning  = "StatusRunning"
	jobStatusDone     = "StatusDone"
	httpMethodGet     = "GET"
	httpMethodPost    = "POST"
	httpMethodPut     = "PUT"
	httpMethodPatch   = "PATCH"
	httpMethodDelete  = "DELETE"
	httpMethodHead    = "HEAD"
	httpMethodOptions = "OPTIONS"
	httpMethodConnect = "CONNECT"
	httpMethodTrace   = "TRACE"
)

func main() {

	possibleStatus := []string{
		jobStatusPending,
		jobStatusRunning,
		jobStatusDone,
	}

	intervalSeconds := rand.IntN(100)
	maxDurationSeconds := rand.IntN(1000)
	jobStatusIndex := rand.IntN(len(possibleStatus))

	jobInstance := Job{
		Name:        "Job_Struct_1",
		Interval:    intervalSeconds,
		MaxDuration: maxDurationSeconds,
		Status:      possibleStatus[jobStatusIndex],
		Task:        HttpTask{httpUrl: "https://stopjava.com", httpMethod: "GEaT"},
	}

	jobStatusStr, jobStatusErr := jobInstance.DescribeJobStatus()
	if jobStatusErr == nil {
		fmt.Println(jobStatusStr)
	} else {
		// fmt.Println("Error: Job status could not be found. JobStatus must be non-empty string")
		fmt.Println(jobStatusErr)
	}

	jobDurationStr, jobDurationErr := jobInstance.DescribeJobDuration()
	if jobDurationErr == nil {
		fmt.Println(jobDurationStr)
	} else {
		// fmt.Println("Error: Job duration could not be found. Interval must be non-zero positive integer")
		fmt.Println(jobDurationErr)
	}

	if err := jobInstance.Task.Run(); err != nil {
		// fmt.Println("Error: PrintTask errored out. debug that")
		fmt.Println(err)
	} else {
		fmt.Printf("I don't need it here since I already print in the implementation but leaving it here for the spirit of it")
	}
	// else {
	// 	fmt.Printf("I don't need it here since I already print in the implementation but leaving it here for the spirit of it")
	// }
}

func (j Job) DescribeJobStatus() (string, error) {
	if j.Status == "" {
		return "", errors.New("Empty Job Status")
	}
	switch j.Status {
	case jobStatusPending:
		return "Job Pending", nil
	case jobStatusRunning:
		return "Job Running", nil
	case jobStatusDone:
		return "Job Done", nil
	default:
		return "", fmt.Errorf("%v is not a valid job status", j.Status)
	}
}

func (j Job) DescribeJobDuration() (string, error) {
	if j.Interval <= 0 || j.Name == "" {
		return "", errors.New("Job Interval or Job Name are empty")
	}
	jobDurationFormattedString := fmt.Sprintf("Job named '%v' ran for '%v' seconds", j.Name, j.Interval)
	return jobDurationFormattedString, nil
}

type Task interface {
	Run() error
}

type PrintTask struct {
	printString string
}

func (pt PrintTask) Run() error {
	if pt.printString == "" {
		return errors.New("Empty Print string")
	}
	fmt.Println(pt.printString)
	return nil
}

type SleepTask struct {
	sleepDurationSec int
}

func (st SleepTask) Run() error {
	if st.sleepDurationSec <= 0 {
		return errors.New("Sleep duration must be positive integer")
	}
	fmt.Printf("Sleeping for %v seconds", st.sleepDurationSec)
	time.Sleep(time.Duration(st.sleepDurationSec) * time.Second)
	return nil
}

type HttpTask struct {
	httpUrl    string
	httpMethod string
}

func (ht HttpTask) Run() error {
	possibleHttpMethods := []string{
		httpMethodGet,
		httpMethodPost,
		httpMethodPut,
		httpMethodPatch,
		httpMethodDelete,
		httpMethodHead,
		httpMethodOptions,
		httpMethodConnect,
		httpMethodTrace,
	}

	if ht.httpUrl == "" || !slices.Contains(possibleHttpMethods, ht.httpMethod) {
		return errors.New("Empty HTTP Task string or Invalid HTTP Method\n")
	}
	fmt.Printf("Performing HTTP %v at URL: %v\n", ht.httpMethod, ht.httpUrl)
	return nil
}
