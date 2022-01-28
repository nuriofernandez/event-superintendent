package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Start server...")

	// Listen on port 8000. 🦻
	ln, _ := net.Listen("tcp", ":8000")
	fmt.Printf("Listening connections at %s!\n", ln.Addr().String())

	for {
		// Accept connection. 🤝
		conn, _ := ln.Accept()

		// Handle the connection async. 🛤
		go handleConection(conn)
	}
}
