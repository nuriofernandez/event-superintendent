package main

import (
	_ "github.com/joho/godotenv/autoload" // Automatically loads godotenv .env file contents. 💖

	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Starting server...")

	// Listen on port {PORT} env variable port. 🦻
	ln, _ := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	fmt.Printf("Listening connections at %s!\n", ln.Addr().String())

	for {
		// Accept connection. 🤝
		conn, _ := ln.Accept()

		// Handle the connection async. 🛤
		go handleConection(conn)
	}
}
