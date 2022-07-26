package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	pb "producer_consumer/pb/message"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr  = flag.String("addr", "localhost", "The address of the manager.")
	port  = flag.String("port", "5011", "The port to the manager.")
	name  = flag.String("name", "producer", "Name of the producer")
	delay = flag.Duration("delay", 1*time.Second, "Delay between processed tasks")
)

func consume_tasks(client pb.ManagerClient) {
	for {
		log.Println("Getting new task:")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		reply, err := client.GetTask(ctx, &pb.GetTaskRequest{ConsumerId: *name})
		cancel()

		if err != nil {
			fail_delay := time.Duration(rand.Intn(1000)) * time.Millisecond
			log.Printf("    - Failed to get task: %s", err.Error())
			log.Printf("    - Waiting %s before next try", fail_delay.String())
			time.Sleep(fail_delay)
		} else {
			log.Printf("    - Processing %s...", reply.GetTaskId())
			time.Sleep(*delay)
			log.Printf("    - %s processed with success!", reply.GetTaskId())
		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", *addr, *port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to reach %s:%s", *addr, *port)
	}
	defer conn.Close()

	client := pb.NewManagerClient(conn)
	consume_tasks(client)
}
