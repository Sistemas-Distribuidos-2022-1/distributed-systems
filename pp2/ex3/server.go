/*
 * Title: Server
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
	fmt.Printf("Listening on %s:%s", SERVER_HOST, SERVER_PORT)

	// Wait for connection
	for {
		connection, err := server.Accept()
		if err != nil {
			panic("Error receiving connection!")
		}

		// Dispatch connection to worker
		go processRequest(connection)
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
	if strings.Compare(message[:3], "APC") == 0 {
		data := strings.Split(message[4:], " ")
		n1, _ := strconv.ParseFloat(data[0], 64)
		n2, _ := strconv.ParseFloat(data[1], 64)
		n3, _ := strconv.ParseFloat(data[2], 64)
		result := "FAIL"

		m := (n1 + n2) / 2.0
		if m >= 7.0 {
			result = "TRUE"
		} else if m <= 3.0 {
			result = "FALSE"
		} else {
			m = (m + n3) / 2.0
			if m >= 5.0 {
				result = "TRUE"
			} else {
				result = "FALSE"
			}
		}

		_, err := conn.Write([]byte(result))
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
