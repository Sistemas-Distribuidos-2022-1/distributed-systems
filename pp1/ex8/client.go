/*
 * Title: Client Ex. 8
 * Author: William T. P. Junior
 * Made with GO 1.18
 * Use: go run client.go
 */
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "50505"
	SERVER_TYPE = "tcp"
)

func main() {
	fmt.Print("-------------------- Special Credit Checker --------------------\n")
	fmt.Print("Element format: AVG_BALANCE\n")
	fmt.Print("Multiple data input: NOT SUPPORTED!\n")
	fmt.Print("type 'exit' to close\n")
	reader := bufio.NewReader(os.Stdin)
	for {
		// Collect input data from terminal
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		input = strings.Replace(input, "\r", "", -1)

		if strings.Compare(input, "exit") == 0 {
			break
		}

		// Passes to the processing function
		special_credit_check(input)
	}
}

func special_credit_check(balance string) {
	// Checks if input value is valid.
	if _, err := strconv.ParseFloat(balance, 64); err != nil {
		panic("Invalid AVG_BALANCE input!")
	}

	// Connect with the server
	conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Send request for net salary
	_, err = conn.Write([]byte(fmt.Sprintf("SCC %s", balance)))
	if err != nil {
		panic("Failed to write to socket!")
	}

	// Receive results
	buffer := make([]byte, 64)
	mLen, err := conn.Read(buffer)
	if err != nil {
		panic("Failed to read from socket!")
	}

	// Present results
	credit, _ := strconv.ParseFloat(string(buffer[:mLen]), 64)
	fmt.Printf("SPECIAL CREDIT AVAILABLE: %.2f\n", credit)
}
