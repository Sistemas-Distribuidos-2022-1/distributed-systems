/*
 * Title: Client Ex. 6
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
	fmt.Print("-------------------- Net Salary Calculator --------------------\n")
	fmt.Print("Element format: NAME LEVEL SALARY NUM_DEPS\n")
	fmt.Print("Multiple data input: NAME1 LEVEL1 SALARY1 NUM_DEPS1 NAME2 LEVEL2 SALARY2 NUM_DEPS2 ...\n")
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
		if len(data)%4 != 0 {
			panic("Invalid input!")
		}

		// Counts the number o employees present in the input data.
		numEmployees := len(data) / 4
		for i := 0; i < numEmployees; i++ {
			// Passes to the processing function
			calcule_net_salary(data[(i*3)], data[(i*3)+1], data[(i*3)+2], data[(i*3)+3])
		}
	}
}

func calcule_net_salary(name, level, salary, num_deps string) {
	// Checks if input value is valid.
	if _, err := strconv.ParseFloat(salary, 64); err != nil {
		panic("Invalid SALARY input!")
	}
	if _, err := strconv.ParseInt(num_deps, 10, 64); err != nil {
		panic("Invalid NUM_DEPS input!")
	}

	// Connect with the server
	conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Send request for net salary
	_, err = conn.Write([]byte(fmt.Sprintf("CNS %s %s %s", strings.ToUpper(level), salary, num_deps)))
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
	net_salary, _ := strconv.ParseFloat(string(buffer[:mLen]), 64)
	fmt.Printf("NAME: %s | LEVEL: %s | NET SALARY: %.2f\n", name, level, net_salary)
}
