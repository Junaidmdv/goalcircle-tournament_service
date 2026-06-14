package server

import (
	"github.com/Junaidmdv/goalcircle-tournament_service/internal/config"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	Server *grpc.Server
	Config *config.Config
}

func NewGRPCServer(cnfg *config.Config) *GRPCServer {

	server := grpc.NewServer()

	return &GRPCServer{
		Server: server,
	}
}

func (gs *GRPCServer) BootstrapSetup() {

}
