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
		Task:        HttpTask{httpUrl: "https://stopjava.com", httpMethod: "GET"},
	}

	jobStatusStr, jobStatusErr := jobInstance.DescribeJobStatus()
	if jobStatusErr == nil {
		fmt.Println(jobStatusStr)
	} else {
		fmt.Println(jobStatusErr)
	}

	jobDurationStr, jobDurationErr := jobInstance.DescribeJobDuration()
	if jobDurationErr == nil {
		fmt.Println(jobDurationStr)
	} else {
		fmt.Println(jobDurationErr)
	}

	if err := jobInstance.Task.Run(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("I don't need it here since I already print in the implementation but leaving it here for the spirit of it\n")
	}
}

func (j Job) DescribeJobStatus() (string, error) {
	if j.Status == "" {
		return "", errors.New("empty job status")
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
	if j.Interval <= 0 {
		return "", fmt.Errorf("job interval %v is invalid. job interval must be a positive integer", j.Interval)
	}
	if j.Name == "" {
		return "", fmt.Errorf("job Name is empty")
	}
	jobDurationFormattedString := fmt.Sprintf("job named '%v' ran for '%v' seconds", j.Name, j.Interval)
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
		return errors.New("empty print string")
	}
	fmt.Println(pt.printString)
	return nil
}

type SleepTask struct {
	sleepDurationSec int
}

func (st SleepTask) Run() error {
	if st.sleepDurationSec <= 0 {
		return errors.New("sleep duration must be positive integer")
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

	if ht.httpUrl == "" {
		return errors.New("empty HTTP task string")
	}
	if !slices.Contains(possibleHttpMethods, ht.httpMethod) {
		return fmt.Errorf("invalid HTTP method %v", ht.httpMethod)
	}
	fmt.Printf("Performing HTTP %v at URL: %v\n", ht.httpMethod, ht.httpUrl)
	return nil
}
