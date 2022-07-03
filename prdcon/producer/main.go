package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "prdcon/pb/message"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = "localhost"
	port = 50051
	name = flag.String("name", "producer", "Name of the producer")
)

func produce_tasks(client pb.TaskManagerClient) {
	counter := 1
	for true {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		_, err := client.RegisterTask(ctx, &pb.RegisterTaskRequest{Id: *name, TaskName: fmt.Sprintf("task_%d", counter)})
		if err != nil {
			log.Fatalf("Could not register task %d", counter)
		} else {
			// log.Printf("Server replied: %s", reply.GetStatus())
			counter += 1
		}
		time.Sleep(3 * time.Second)
	}
}

func main() {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", addr, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to reach %s:%d", addr, port)
	}
	defer conn.Close()
	client := pb.NewTaskManagerClient(conn)

	produce_tasks(client)
}
