package main

import (
	"context"
	"fmt"
	"log"
	"net"

	// This import path is based on the name declaration in the go.mod,
	// and the gen/proto/go output location in the buf.gen.yaml.
	echo "github.com/awe76/sagadt/echoapis/v1"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listenOn := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}

	server := grpc.NewServer()
	echo.RegisterEchoServiceServer(server, &echoServiceServer{})
	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

// echoServiceServer implements the EchoService API.
type echoServiceServer struct {
	echo.UnimplementedEchoServiceServer
}

// Echo sends input value backward.
func (s *echoServiceServer) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {

	value := req.GetValue()
	log.Println("Echo: ", value)

	return &echo.EchoResponse{Value: value}, nil
}
