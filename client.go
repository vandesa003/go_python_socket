package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
)

func main() {
	// Start the client and connect to the server.
	fmt.Println("Connecting to " + connType + " server " + connHost + ":" + connPort)
	conn, err := net.Dial(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	// Create new reader from Stdin.
	reader := bufio.NewReader(os.Stdin)

	// run loop forever, until exit.
	for {
		// Prompting message.
		fmt.Print("Test to send: ")

		// Read in input until newline, Enter key.
		input, _ := reader.ReadString('\n')

		// Send to socket connection.
		conn.Write([]byte(input))

		// Listen for relay.
		message, _ := bufio.NewReader(conn).ReadString('\n')

		// Print server relay.
		log.Print("Server Reply: ", message)
	}
}
