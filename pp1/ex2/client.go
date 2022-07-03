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
	fmt.Print("-------------------- Majority Age Check --------------------\n")
	fmt.Print("Element format: NAME GENDER(M/F) AGE\n")
	fmt.Print("Multiple data input: NAME1 GENDER1 AGE1 NAME2 GENDER2 AGE2 ...\n")
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
		if len(data)%3 != 0 {
			panic("Invalid input!")
		}

		// Counts the number o peoples present in the input data.
		numPeoples := len(data) / 3
		for i := 0; i < numPeoples; i++ {
			// Checks if age value is valid.
			if _, err := strconv.ParseInt(data[(i*3)+2], 10, 64); err != nil {
				panic("Invalid input!")
			}

			// Passes to the processing function
			majority(data[(i*3)], data[(i*3)+1], data[(i*3)+2])
		}
	}
}

func majority(name, gender, age string) {
	// Connect with the server
	conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Send request for age check
	_, err = conn.Write([]byte(fmt.Sprintf("MAJ %s %s %s", name, strings.ToUpper(gender), age)))
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
		fmt.Printf("%s has reached the age of majority.\n", name)
	} else if strings.Compare(string(buffer[:mLen]), "FALSE") == 0 {
		fmt.Printf("%s has not reached the age majority.\n", name)
	} else {
		fmt.Printf("It was not possible to check the age of majority for {%s, %s, %s}\n", name, gender, age)
	}
}
