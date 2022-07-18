package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	// Notify new connections. 🔊
	ip := conn.RemoteAddr().String()
	fmt.Printf("<%s> Has been connected!\n", ip)

	// Process data until the end of the process. 🔁
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')

		// Detect network errors and close connection. 👋
		if err != nil {
			fmt.Printf("<%s> Has been disconnected!\n", ip)
			conn.Close()
			return
		}

		// Log the received message. 🔊
		logMsg := string(message[:len(message)-2]) // Remove \n from the string
		fmt.Printf("<%s> : %s\n", ip, logMsg)
	}
}
