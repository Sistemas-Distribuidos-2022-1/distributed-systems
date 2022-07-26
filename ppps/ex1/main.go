package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	goczmq "github.com/zeromq/goczmq"
)

const (
	PUB_ADDR  = "tcp://*:5012"
	PUB_GROUP = "Q1"
)

func main() {
	pub, err := goczmq.NewPub(PUB_ADDR)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer pub.Destroy()

	go produce("John", "operador", "100", pub)
	produce("Anna", "programador", "100", pub)
}

func produce(name, role, salary string, pub *goczmq.Sock) {
	for {
		message := [][]byte{}
		message = append(message, []byte(PUB_GROUP))
		message = append(message, []byte("\033[32m"+name+"\033[0m"))
		message = append(message, []byte(role))
		message = append(message, []byte(salary))

		delay := time.Duration(rand.Intn(1000)) * time.Millisecond
		err := pub.SendMessage(message)
		if err != nil {
			log.Println("\033[31m" + err.Error() + "\033[0m")
			log.Printf("\033[31m"+"Waiting %s before next try"+"\033[0m", delay)
		} else {
			log.Println("\033[32m" + "Message sent with success!" + "\033[0m")
			salary_val, _ := strconv.Atoi(salary)
			salary = fmt.Sprintf("%d", salary_val+1)
		}

		time.Sleep(delay)
	}
}
