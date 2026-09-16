package main

import (
	"fmt"
)

func main() {
	var interval_seconds int
	var job_status string
	name := "Worker"

	string, _ := fmt.Printf("Worker name: '%v' at Status: '%v' after Interval: '%v'\n", name, job_status, interval_seconds)
	fmt.Println(string)
}
