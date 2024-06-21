package erd

import (
	"MEZ/logger"
	"fmt"
	"net"
)

func TrapClient() {
	serverAddress, err := net.ResolveUDPAddr("udp", ":162")
	if err != nil {
		logger.Log("ResolveUDPAddr:", err)
		return
	}
	connection, err := net.ListenUDP("udp", serverAddress)
	if err != nil {
		logger.Log("ListenUDP:", err)
		return
	}
	defer connection.Close()

	for {
		inputBytes := make([]byte, 1024)
		connection.ReadFromUDP(inputBytes)
		fmt.Println(string(inputBytes))
	}
}
