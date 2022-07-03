/*
 * Title: Client
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
	fmt.Print("-------------------- Grade Check --------------------\n")
	fmt.Print("Element format: N1 N2 N3\n")
	fmt.Print("Multiple data input: NOT SUPPORTED!\n")
	fmt.Print("type 'exit' to close\n")
	reader := bufio.NewReader(os.Stdin)
	for {
		// Collect input data from terminal
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)

		if strings.Compare(input, "exit") == 0 {
			break
		}

		// Check that the input string has the correct format
		data := strings.Split(input, " ")
		if len(data) != 3 {
			panic("Invalid input!")
		}

		// Passes to the processing function
		approvalCheck(data[0], data[1], data[2])

	}
}

func approvalCheck(n1, n2, n3 string) {
	// Checks if grade values are valid.
	if _, err := strconv.ParseFloat(n1, 64); err != nil {
		panic("Invalid N1 input!")
	}
	if _, err := strconv.ParseFloat(n2, 64); err != nil {
		panic("Invalid N2 input!")
	}
	if _, err := strconv.ParseFloat(n3, 64); err != nil {
		panic("Invalid N3 input!")
	}

	// Connect with the server
	conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Send request for approval check
	_, err = conn.Write([]byte(fmt.Sprintf("APC %s %s %s", n1, n2, n3)))
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
	if strings.Compare(string(buffer[:mLen]), "TRUE") == 0 {
		fmt.Printf("Approved!\n")
	} else if strings.Compare(string(buffer[:mLen]), "FALSE") == 0 {
		fmt.Printf("Reproved!\n")
	} else {
		fmt.Printf("It was not possible to check approval for grades {%s, %s, %s}\n", n1, n2, n3)
	}
}
