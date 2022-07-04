/*
 * Title: Client Ex. 5
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
	fmt.Print("-------------------- Swimmer Age Category --------------------\n")
	fmt.Print("Element format: AGE\n")
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
		calcule_category(input)
	}
}

func calcule_category(age string) {
	// Checks if age value is valid.
	if _, err := strconv.ParseInt(age, 10, 64); err != nil {
		panic("Invalid Age input!")
	}

	// Connect with the server
	conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Send request for age category
	_, err = conn.Write([]byte(fmt.Sprintf("SAC %s", age)))
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
	fmt.Printf("Category: %s\n", string(buffer[:mLen]))
}
