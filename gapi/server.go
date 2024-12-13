package gapi

import (
	"github.com/ebukacodes21/peerbill-user-api/config"
	db "github.com/ebukacodes21/peerbill-user-api/db/sqlc"
	"github.com/ebukacodes21/peerbill-user-api/pb"
)

type GServer struct {
	pb.UnimplementedPeerbillUserServer
	config     config.Config
	repository db.DatabaseContract
}

func NewGServer(config config.Config, repository db.DatabaseContract) *GServer {
	return &GServer{config: config, repository: repository}
}
