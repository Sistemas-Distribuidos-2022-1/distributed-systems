/*
 * Title: Salary Readjustment Client
 * Description: This client collects the name, role and actual salary
 *              of an employee, then sends to the server that apply a
 *              readjustment.
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
	fmt.Print("-------------------- Employee Salary Readjustment --------------------\n")
	fmt.Print("Element format: NAME ROLE SALARY\n")
	fmt.Print("Multiple data input: NAME1 ROLE1 SALARY1 NAME2 ROLE2 SALARY2 ...\n")
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

		// Counts the number o employees present in the input data.
		numEmployees := len(data) / 3
		for i := 0; i < numEmployees; i++ {
			// Checks if salary value is valid.
			if _, err := strconv.ParseFloat(data[(i*3)+2], 64); err != nil {
				panic("Invalid input!")
			}

			// Passes to the processing function
			readjust(data[(i*3)], data[(i*3)+1], data[(i*3)+2])
		}
	}
}

func readjust(name, role, salary string) {
	// Connect with the server
	conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Send request for readjustment
	_, err = conn.Write([]byte(fmt.Sprintf("RAJ %s %s %s", name, role, salary)))
	if err != nil {
		panic("Failed to write to socket!")
	}

	// Receive results
	buffer := make([]byte, 8)
	mLen, err := conn.Read(buffer)
	if err != nil {
		panic("Failed to read from socket!")
	}

	// Present results
	oldSal, _ := strconv.ParseFloat(salary, 64)
	newSal, _ := strconv.ParseFloat(string(buffer[:mLen]), 64)
	fmt.Printf("NAME: %s | ROLE: %s | SALARY: %.2f -> %.2f\n", name, role, oldSal, newSal)
}
