package main

import (
	"context"
	"fmt"
	"log"
	"net"

	// This import path is based on the name declaration in the go.mod,
	// and the gen/proto/go output location in the buf.gen.yaml.
	petv1 "github.com/alehechka/buf-tour/petstore/gen/proto/go/pet/v1"
	"google.golang.org/genproto/googleapis/type/datetime"
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
	petv1.RegisterPetStoreServiceServer(server, &petStoreServiceServer{})
	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

// petStoreServiceServer implements the PetStoreService API.
type petStoreServiceServer struct {
	petv1.UnimplementedPetStoreServiceServer
}

// PutPet adds the pet associated with the given request into the PetStore.
func (s *petStoreServiceServer) PutPet(ctx context.Context, req *petv1.PutPetRequest) (*petv1.PutPetResponse, error) {
	name := req.GetName()
	petType := req.GetPetType()
	log.Println("Got a request to create a", petType, "named", name)

	return &petv1.PutPetResponse{Pet: &petv1.Pet{PetType: petType, Name: name, PetId: "1234", CreatedAt: &datetime.DateTime{}}}, nil
}
