package db_stream

import (
	"context"
	dbpb "github.com/satk124/microservice/proto/gen/db"
)

type DBStream struct {
}

func New() *dbpb.DbServer {
	var x dbpb.DbServer
	x = DBStream{}
	return &DBStream{}
}

func (dbs DBStream) Get(c context.Context, r *dbpb.Request) dbpb.Response {

}

// func (d DBStream)

// type DbClient interface {
// 	Get(ctx context.Context, in *Request, opts ...grpc.CallOption) (Db_GetClient, error)
// }
