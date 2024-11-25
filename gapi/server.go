package gapi

import (
	"peerbill-user-api/config"
	"peerbill-user-api/pb"
)

type GServer struct {
	pb.UnimplementedPeerbillUserServer
	config config.Config
}

func NewGServer(config config.Config) *GServer {
	return &GServer{config: config}
}
