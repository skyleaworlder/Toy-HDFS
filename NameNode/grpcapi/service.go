package grpcapi

import (
	"log"
	"net"

	"github.com/skyleaworlder/Toy-HDFS/proto"
	"google.golang.org/grpc"
)

// NameNodeServeClient is a function that need client to call for service
func NameNodeServeClient(lsnClient *net.TCPListener) {
	// grpc server
	clientGrpcServer := grpc.NewServer()

	// construct some Servers
	// clientNameNodeServer process requests from client
	clientNameNodeServer := NewClientNameNodeServer()
	proto.RegisterClientNameNodeServer(clientGrpcServer, clientNameNodeServer)

	if err := clientGrpcServer.Serve(lsnClient); err != nil {
		log.Fatal("NameNode.interface.service.go->NameNodeServe error:", err.Error())
	}
}
