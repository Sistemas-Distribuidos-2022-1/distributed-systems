/*
 * Title: Client Ex. 4
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
	fmt.Print("-------------------- Ideal Weight Calculator --------------------\n")
	fmt.Print("Element format: HEIGHT GENDER(M/F)\n")
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
		calcule_weight(data[0], data[1])
	}
}

func calcule_weight(height, gender string) {
	// Checks if height value is valid.
	if _, err := strconv.ParseFloat(height, 64); err != nil {
		panic("Invalid HEIGHT input!")
	}

	// Connect with the server
	conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Send request for ideal weight
	_, err = conn.Write([]byte(fmt.Sprintf("IWC %s %s", height, strings.ToUpper(gender))))
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
	weight, _ := strconv.ParseFloat(string(buffer[:mLen]), 64)
	fmt.Printf("Ideal Weight: %.2f\n", weight)
}
