package main

import (
	"fmt"
	"net"
	"os"
)

// CheckServerError from UDP connection and UDP address
func CheckServerError(err error) {
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}
}

func main() {

	ServerAddr, err := net.ResolveUDPAddr("udp", "localhost:8125")
	CheckServerError(err)
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	// ServerConn, err := net.Listener("udp", ServerAddr)
	CheckServerError(err)

	defer ServerConn.Close()

	buf := make([]byte, 2048)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		fmt.Printf("Received message: %s from  %s\n", string(buf[0:n]), addr)
		CheckServerError(err)
	}

}
