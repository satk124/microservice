package autocomplete

import (
	"context"
	autocompletepb "github.com/satk124/microservice/proto/gen/autocomplete"
)

type Autocomplete struct {
}

func New() Autocomplete {
	return Autocomplete{}
}

func (a Autocomplete) Get(context.Context, *autocompletepb.Request) (*autocompletepb.Response, error) {
	result := []string{
		"hello",
		"hii",
	}
	return &autocompletepb.Response{Result: result}, nil
}
