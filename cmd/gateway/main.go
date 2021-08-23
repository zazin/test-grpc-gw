package main

import (
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	helloworldpb "test-grpc-gw/proto/helloworld"
	location_pb "test-grpc-gw/proto/location"
)

type server struct {
	helloworldpb.UnimplementedGreeterServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	return &helloworldpb.HelloReply{Message: in.Name + " world"}, nil
}

var (
	helloWorldEndpoint = flag.String("hello_endpoint", "localhost:5001", "endpoint of HelloServer")
	locationEndpoint   = flag.String("location_endpoint", "localhost:5002", "endpoint of LocationServer")
)

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	helloworldpb.RegisterGreeterServer(s, &server{})
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8090")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests

	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err = helloworldpb.RegisterGreeterHandlerFromEndpoint(context.Background(), gwmux, *helloWorldEndpoint, opts)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	err = location_pb.RegisterLocationHandlerFromEndpoint(context.Background(), gwmux, *locationEndpoint, opts)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8001",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8001")
	log.Fatalln(gwServer.ListenAndServe())
}
