package sendmetrics

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

// CheckError from UDP connection and UDP address
func CheckError(err error, from string) {
	if err != nil {
		fmt.Println(from, " error: ", err.Error())
		os.Exit(1)
	}
}

// SendMetrics function send metrics to statsd server
func SendMetrics(address string, metricName string, metricNum int, timeout int64, totalCount int) {

	udpAddr, err := net.ResolveUDPAddr("udp", address)
	resUDPAddr := "ResolveUDPAddr"
	CheckError(err, resUDPAddr)

	conn, err := net.DialUDP("udp", nil, udpAddr)
	dialUDP := "DialUDP"
	CheckError(err, dialUDP)

	defer conn.Close()

	w := bufio.NewWriter(conn)

	fmt.Printf("Start generator: %s\n", time.Now().Local().Format("2006-01-02 15:04:05"))

	for time.Now().Unix() <= timeout {
		for i := 0; i <= metricNum; i++ {
			fmt.Fprintf(w, "%s%d.count:%d|c\n", metricName, i, 1)
			totalCount++
			w.Flush()
		}
		time.Sleep(time.Duration(3) * time.Nanosecond)
	}

	fmt.Printf("Ended generator: %s\n", time.Now().Local().Format("2006-01-02 15:04:05"))

	fmt.Printf("Send metrics to statsd server %s, total metrics: %d\n", udpAddr, totalCount)

}
