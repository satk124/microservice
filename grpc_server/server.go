package grpcserver

import (
	// "errors"
	"net"

	autocomplete "github.com/satk124/microservice/autocomplete"
	autocompletepb "github.com/satk124/microservice/proto/gen/autocomplete"
	grpc "google.golang.org/grpc"
)

type Server struct {
}

func New() *Server {
	return &Server{}
}

func (s *Server) Serve() error {
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	autocompletepb.RegisterAutocompleteServer(grpcServer, autocomplete.New())
	listner, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err.Error())
	}
	err = grpcServer.Serve(listner)
	if err != nil {
		panic(err.Error())
	}
	return nil
}
