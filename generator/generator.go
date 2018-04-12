package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {

	address := os.Args[1]
	// port := os.Args[2]
	metricName := os.Args[2]
	metricNum, _ := strconv.Atoi(os.Args[3])
	runTime, _ := strconv.Atoi(os.Args[4])
	// now := time.Now().Unix()
	timeout := time.Now().Unix() + int64(runTime)

	sendMetrics(address, metricName, metricNum, timeout)

}

func sendMetrics(address string, metricName string, metricNum int, timeout int64) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	// conn.Write()
	w := bufio.NewWriter(conn)
	for time.Now().Unix() <= timeout {
		for i := 0; i <= metricNum; i++ {
			fmt.Fprintf(w, "%s.%s.counter:1|c\n", metricName, i)
			// fmt.Println(i)
		}
	}

	// fmt.Printf("%T", port)
	fmt.Println(tcpAddr)
	// fmt.Println(metricNum)
	// fmt.Println(metricName)
	// fmt.Println(now)

	// fmt.Println(runTime)
	// fmt.Println(timeout)
}
