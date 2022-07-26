/*
 * Title: Producer
 * Description: Generates workload over time.
 * Author: William T. P. Junior
 * Made with GO 1.18
 */

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	pb "producer_consumer/pb/message"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr  = flag.String("addr", "localhost", "The address of the manager.")
	port  = flag.String("port", "5011", "The port to the manager.")
	name  = flag.String("name", "producer", "Name of the producer")
	delay = flag.Duration("delay", 1*time.Second, "Delay between generated tasks")
)

/*
 * Produce and send tasks to the manager every delay interval.
 */
func produce_tasks(client pb.ManagerClient, delay time.Duration) {
	counter := 1
	for {
		task_id := fmt.Sprintf("w_%s-%d", *name, counter)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		_, err := client.RegisterTask(ctx, &pb.RegisterTaskRequest{TaskId: task_id, TaskName: "workload", ProducerId: *name})
		cancel()

		if err != nil {
			fail_delay := time.Duration(rand.Intn(1000)) * time.Millisecond
			log.Printf("Failed to register task: %s", err.Error())
			log.Printf("Waiting %s before next try", fail_delay.String())
			time.Sleep(fail_delay)
		} else {
			log.Printf("Registered task: %s", task_id)
			counter += 1
			time.Sleep(delay)
		}
	}
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", *addr, *port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to reach %s:%s", *addr, *port)
	}
	defer conn.Close()

	client := pb.NewManagerClient(conn)
	produce_tasks(client, *delay)
}
