package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Start server...")

	// Listen on port 8000. 🦻
	ln, _ := net.Listen("tcp", ":8000")
	fmt.Printf("Listening connections at %s!\n", ln.Addr().String())

	// Accept connection. 🤝
	conn, _ := ln.Accept()

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
