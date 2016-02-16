package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "datamodel/protobuf"
)

type serverStruct struct {
	g *grpc.Server
}

var server *serverStruct

// Run ...
func Run(port uint) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port)) // RPC port
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	g := grpc.NewServer()

	server = &serverStruct{g}
	pb.RegisterSkizzeServer(g, server)
	_ = g.Serve(lis)
}
