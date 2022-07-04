/*
 * Title: Server Ex. 6
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
	if strings.Compare(message[:3], "CNS") == 0 {
		data := strings.Split(message[4:], " ")
		salary, _ := strconv.ParseFloat(data[1], 64)
		num_deps, _ := strconv.ParseInt(data[2], 10, 64)

		result := -1.0
		switch data[0] {
		case "A":
			if num_deps > 0 {
				result = salary * 0.92
			} else {
				result = salary * 0.97
			}
		case "B":
			if num_deps > 0 {
				result = salary * 0.90
			} else {
				result = salary * 0.95
			}
		case "C":
			if num_deps > 0 {
				result = salary * 0.85
			} else {
				result = salary * 0.92
			}
		case "D":
			if num_deps > 0 {
				result = salary * 0.83
			} else {
				result = salary * 0.90
			}
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
