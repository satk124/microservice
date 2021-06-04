package main

import (
	"context"
	"fmt"

	autocompletepb "github.com/satk124/microservice/proto/gen/autocomplete"
	"google.golang.org/grpc"
)

func main() {
	serverAddr := "localhost:8081"
	c, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}

	client := autocompletepb.NewAutocompleteClient(c)
	resp, err := client.Get(context.TODO(), &autocompletepb.Request{Key: "hii", Size: 5})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%v, %T", resp, resp)
}
