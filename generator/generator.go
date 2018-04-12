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
	metricName := os.Args[2]
	metricNum, _ := strconv.Atoi(os.Args[3])
	runTime, _ := strconv.Atoi(os.Args[4])
	// now := time.Now().Unix()
	timeout := time.Now().Unix() + int64(runTime)

	sendMetrics(address, metricName, metricNum, timeout)

}

func sendMetrics(address string, metricName string, metricNum int, timeout int64) {

	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		println("ResolveUDPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	w := bufio.NewWriter(conn)
	// fmt.Fprint(w, "Hello, ")

	for time.Now().Unix() <= timeout {
		for i := 0; i <= metricNum; i++ {
			fmt.Fprintf(w, "%s--%s.count:%d|c\n", "proba", metricName, 1)
			w.Flush()
		}
	}

	// fmt.Printf("%T", port)
	fmt.Println(udpAddr)
	// fmt.Println(conn)
	// fmt.Println(metricNum)
	// fmt.Println(metricName)
	// fmt.Println(now)

	// fmt.Println(runTime)
	// fmt.Println(timeout)
}
