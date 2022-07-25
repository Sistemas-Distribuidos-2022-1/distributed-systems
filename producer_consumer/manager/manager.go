/*
 * Title: Manager
 * Description: Provides intermediary services to a producer consumer system.
 * Author: William T. P. Junior
 * Made with GO 1.18
 */

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	bf "producer_consumer/manager/buffer"
	pb "producer_consumer/pb/message"

	"github.com/gosuri/uilive"
	"google.golang.org/grpc"
)

var port = 5011

type server struct {
	pb.UnimplementedManagerServer
	buffer bf.Buffer
}

/*
 * RPC RegisterTask Service
 */
func (s *server) RegisterTask(ctx context.Context, in *pb.RegisterTaskRequest) (*pb.RegisterTaskReply, error) {
	// Try to add the received task to the buffer
	err := s.buffer.Add(bf.Task{Id: in.TaskId, Name: in.TaskName, Producer: in.ProducerId})
	if err != nil {
		return &pb.RegisterTaskReply{TaskId: in.TaskId, Success: false}, err
	}
	return &pb.RegisterTaskReply{TaskId: in.TaskId, Success: true}, nil
}

/*
 * RPC GetTask Service
 */
func (s *server) GetTask(ctx context.Context, in *pb.GetTaskRequest) (*pb.GetTaskReply, error) {
	// Try to get a task from the buffer
	task, err := s.buffer.Pop()
	if err != nil {
		return &pb.GetTaskReply{TaskId: "?", TaskName: "?", Success: false}, err
	}
	return &pb.GetTaskReply{TaskId: task.Id, TaskName: task.Name, Success: true}, nil
}

/*
 * Prints and updates the task list with delay interval.
 */
func (s *server) print_task_list(delay time.Duration) {
	// Create live writer
	writer := uilive.New()
	writer.Start()
	defer writer.Stop()

	for true {
		text := "------ PENDING TASK LIST ------\n"

		// Get the task list
		task_list := s.buffer.ListTasks()
		if task_list != nil {
			for i := range task_list {
				text += ("* " + task_list[i].Id + " - " + task_list[i].Name + "\n")
			}
		}

		text += "-------------------------------"

		//Update the live writer
		fmt.Fprintln(writer, text)

		// Wait for next update time
		time.Sleep(delay * time.Second)
	}
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	svc := grpc.NewServer()
	svr := &server{}
	pb.RegisterManagerServer(svc, svr)

	log.Printf("Started server on port %d", port)
	go svr.print_task_list(1)

	if err := svc.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
