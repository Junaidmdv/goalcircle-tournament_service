package server

import "google.golang.org/grpc"

type GRPCServer struct {
	Server *grpc.Server
}




func NewGRPCServer() *GRPCServer {

	server := grpc.NewServer()

	return &GRPCServer{
		Server: server,
	}
}







func(gs *GRPCServer)BootstrapSetup(){

}