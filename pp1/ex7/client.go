/*
 * Title: Client Ex. 7
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
	fmt.Print("-------------------- Retirement Availability Calculator --------------------\n")
	fmt.Print("Element format: AGE SERVICE_TIME\n")
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

		// Check that the input string has the correct format
		data := strings.Split(input, " ")
		if len(data) != 2 {
			panic("Invalid input!")
		}

		// Passes to the processing function
		retirement_check(data[0], data[1])
	}
}

func retirement_check(age, service string) {
	// Checks if input value is valid.
	if _, err := strconv.ParseInt(age, 10, 64); err != nil {
		panic("Invalid AGE input!")
	}
	if _, err := strconv.ParseInt(service, 10, 64); err != nil {
		panic("Invalid SERVICE_TIME input!")
	}

	// Connect with the server
	conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Send request for net salary
	_, err = conn.Write([]byte(fmt.Sprintf("RAC %s %s", age, service)))
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
		fmt.Printf("Available!\n")
	} else if strings.Compare(string(buffer[:mLen]), "FALSE") == 0 {
		fmt.Printf("Unavailable!\n")
	} else {
		fmt.Printf("It was not possible to check retirement availability for {%s, %s}\n", age, service)
	}
}
