package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	pb "prdcon/pb/message"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = "localhost"
	port = 50051
	name = flag.String("name", "producer", "Name of the producer")
)

func consume_tasks(client pb.TaskManagerClient) {
	for true {
		log.Println("Getting new task!")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		reply, err := client.GetTask(ctx, &pb.GetTaskRequest{Id: *name})
		if err != nil {
			log.Fatalf("Failed to get task from task manager")
		}
		if reply.GetStatus() == 1 {
			log.Printf("Processing %s...", reply.GetTaskName())
			time.Sleep(4 * time.Second)
			log.Printf("%s processed with success!", reply.GetTaskName())
		} else {
			log.Println("No tasks available.")
			time.Sleep(4 * time.Second)
		}
	}
}

func main() {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", addr, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to reach %s:%d", addr, port)
	}
	defer conn.Close()
	client := pb.NewTaskManagerClient(conn)

	consume_tasks(client)
}
