package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/ysytyuk/golang/generator/sendmetrics"
)

func main() {

	if len(os.Args) == 1 || len(os.Args) < 5 || len(os.Args) > 5 || (os.Args[1] == "-h" || os.Args[1] == "--help") {
		fmt.Printf("Help for statsd generator \nfirst argument destination address with port(example: localhost:8125), second metric name, third number metrics, last time running script per seconds\n")
		os.Exit(0)
	}

	address := os.Args[1]
	metricName := os.Args[2]
	metricNum, _ := strconv.Atoi(os.Args[3])
	runTime, _ := strconv.Atoi(os.Args[4])
	// now := time.Now().Unix()
	timeout := time.Now().Unix() + int64(runTime)
	var totalCount int

	sendmetrics.SendMetrics(address, metricName, metricNum, timeout, totalCount)

}
