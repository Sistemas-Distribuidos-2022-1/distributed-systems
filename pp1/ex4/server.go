/*
 * Title: Server Ex. 4
 * Author: William T. P. Junior
 * Made with GO 1.18
 * Use: go run server.go
 */
package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "50505"
	SERVER_TYPE = "tcp"
)

func main() {
	// Getting server running
	fmt.Println("Starting Server...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic("Failed to listen on port" + SERVER_PORT)
	}
	defer server.Close()
	fmt.Printf("Listening on %s:%s\n", SERVER_HOST, SERVER_PORT)

	// Wait for connection
	for {
		connection, err := server.Accept()
		if err != nil {
			panic("Error receiving connection!")
		}

		// Process connection
		processRequest(connection)
	}
}

// Worker function
func processRequest(conn net.Conn) {
	// Read from socket to buffer
	buffer := make([]byte, 1024)
	mLen, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("Failed to read from socket!\n")
		conn.Close()
		return
	}

	// Check if message is valid then treat it
	message := string(buffer[:mLen])
	if strings.Compare(message[:3], "IWC") == 0 {
		data := strings.Split(message[4:], " ")
		height, _ := strconv.ParseFloat(data[0], 64)

		result := -1.0
		if strings.Compare(data[1], "M") == 0 {
			result = (72.7 * height) - 58.0
		} else if strings.Compare(data[1], "F") == 0 {
			result = (62.1 * height) - 44.7
		}

		_, err := conn.Write([]byte(fmt.Sprintf("%f", result)))
		if err != nil {
			fmt.Printf("Failed to write to socket!\n")
			conn.Close()
			return
		}
	} else {
		fmt.Printf("Invalid message!")
	}
	conn.Close()
}
