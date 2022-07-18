package main

import (
	"github.com/xXNurioXx/event-superintendent/v2/connections"

	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	// Notify new connections. 🔊
	ip := conn.RemoteAddr().String()
	fmt.Printf("<%s> Has been connected!\n", ip)
	connection := connections.AddConnection(conn)

	// Process data until the end of the process. 🔁
	for {
		message, err := connection.Reader.ReadString('\n')

		// Detect network errors and close connection. 👋
		if err != nil {
			fmt.Printf("<%s> Has been disconnected!\n", ip)
			connections.RemoveConnection(connection)
			conn.Close()
			return
		}

		// Log the received message. 🔊
		logMsg := string(message[:len(message)-2]) // Remove \n from the string
		fmt.Printf("<%s> : %s\n", ip, logMsg)
	}
}
