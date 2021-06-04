package main

import (
	"fmt"
	grpc_server "github.com/satk124/microservice/grpc_server"
)

func main() {
	fmt.Println("hii")
	g := grpc_server.New()
	g.Serve()
}
