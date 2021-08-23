package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	country "test-grpc-gw/internal/location"
	"test-grpc-gw/internal/location/models"
	location_pb "test-grpc-gw/proto/location"
)

type server struct{}

func countryGrpc(c models.Country) *location_pb.Country {
	return &location_pb.Country{
		Id:   int32(c.ID),
		Code: c.Code,
		Name: c.Name,
	}
}

func (s server) FindCountry(ctx context.Context, empty *location_pb.Empty) (*location_pb.CountryList, error) {
	cnt := country.New()
	countries := cnt.GetCountries()

	items := []*location_pb.Country{}
	for _, c := range countries {
		items = append(items, countryGrpc(c))
	}

	return &location_pb.CountryList{Countries: items}, nil
}

func NewServer() *server {
	return &server{}
}

func main() {

	lis, err := net.Listen("tcp", ":5002")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	location_pb.RegisterLocationServer(s, &server{})
	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:5002")
	log.Fatal(s.Serve(lis))
}
