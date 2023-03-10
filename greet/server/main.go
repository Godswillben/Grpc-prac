package main

import (
	"log"
	"net"

	pb "github.com/Godswillben/go-grpc-practice/greet/proto"
	"google.golang.org/grpc"
)

var addr = "0.0.0.0:50051";

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err !=  nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listenting on %s \n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n, err")
	}
}
