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
	}

	// fmt.Println(jobInstance)

	jobStatusStr, jobStatusBool := jobInstance.DescribeJobStatus()
	if jobStatusBool {
		fmt.Println(jobStatusStr)
	} else {
		fmt.Println("Error: Job status could not be found. JobStatus must be non-empty string")
	}

	jobDurationStr, jobDurationBool := jobInstance.DescribeJobDuration()
	if jobDurationBool {
		fmt.Println(jobDurationStr)
	} else {
		fmt.Println("Error: Job duration could not be found. Interval must be non-zero positive integer")
	}
}

func (j Job) DescribeJobStatus() (string, bool) {
	if j.Status == "" {
		return "", false
	}
	switch j.Status {
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

func (j Job) DescribeJobDuration() (string, bool) {
	if j.Interval <= 0 || j.Name == "" {
		return "", false
	}
	jobDurationFormattedString := fmt.Sprintf("Job named '%v' ran for '%v' seconds", j.Name, j.Interval)
	return jobDurationFormattedString, true
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
		return errors.New("Empty HTTP Task string or Invalid HTTP Method")
	}
	fmt.Printf("Performing HTTP %v at URL: %v", ht.httpMethod, ht.httpUrl)
	return nil
}
