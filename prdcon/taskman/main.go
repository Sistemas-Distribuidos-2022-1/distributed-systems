package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "prdcon/pb/message"

	"github.com/gosuri/uilive"
	"google.golang.org/grpc"
)

var port = 50051

type server struct {
	pb.UnimplementedTaskManagerServer
	queue []string
}

func (s *server) RegisterTask(ctx context.Context, in *pb.RegisterTaskRequest) (*pb.RegisterTaskReply, error) {
	// log.Printf("Received task %s", in.GetTaskName())
	if s.queue == nil {
		s.queue = make([]string, 0)
	}
	s.queue = append(s.queue, fmt.Sprintf("%s_%s", in.GetId(), in.GetTaskName()))
	return &pb.RegisterTaskReply{Status: "Task " + in.GetTaskName() + " received!"}, nil
}

func (s *server) GetTask(ctx context.Context, in *pb.GetTaskRequest) (*pb.GetTaskReply, error) {
	if s.queue == nil {
		s.queue = make([]string, 0)
	}
	if len(s.queue) == 0 {
		return &pb.GetTaskReply{Status: 0, TaskName: ""}, nil
	}
	task := s.queue[0]
	s.queue = s.queue[1:]
	return &pb.GetTaskReply{Status: 1, TaskName: task}, nil
}

func (s *server) print_task_list() {
	writer := uilive.New()
	writer.Start()
	defer writer.Stop()
	for true {
		tasks := "------ PENDING TASK LIST ------\n"
		if s.queue != nil {
			for task := range s.queue {
				tasks += "* " + s.queue[task] + "\n"
			}
		}
		tasks += "-------------------------------"
		fmt.Fprintln(writer, tasks)
		time.Sleep(1 * time.Second)
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
	pb.RegisterTaskManagerServer(svc, svr)

	log.Printf("Started server on port %d", port)
	go svr.print_task_list()
	if err := svc.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
