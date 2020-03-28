package main

import (
	"github.com/pytorchtw/grpc-hello/go_server/controller/hello_controller"
	"github.com/pytorchtw/grpc-hello/go_server/proto/hello"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	Address = "0.0.0.0:9090"
)

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	hello.RegisterHelloServer(s, &hello_controller.HelloController{})

	log.Println("Listen on " + Address)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
