package grpcapi

import (
	"log"
	"net"

	"github.com/skyleaworlder/Toy-HDFS/proto"
	"google.golang.org/grpc"
)

// NameNodeServeClient is a function that need namenode to call for service
func NameNodeServeClient(lsnClient *net.TCPListener) {
	// grpc server
	clientGrpcServer := grpc.NewServer()

	// construct some Servers
	// clientNameNodeServer process requests from client
	clientNameNodeServer := NewClientNameNodeServer()
	proto.RegisterClientNameNodeServer(clientGrpcServer, clientNameNodeServer)

	if err := clientGrpcServer.Serve(lsnClient); err != nil {
		log.Fatal("NameNode.grpcapi.service.go->NameNodeServe error:", err.Error())
	}
}

// NameNodeServeDataNode is a function that need namenode to call for service
func NameNodeServeDataNode(lsnDataNode *net.TCPListener) {
	// grpc server
	datanodeGrpcServer := grpc.NewServer()

	// construct some Servers
	// datanodeNameNodeServer process requests from datanode
	datanodeNameNodeServer := NewDataNodeNameNodeServer()
	proto.RegisterDataNodeNameNodeServer(datanodeGrpcServer, datanodeNameNodeServer)

	if err := datanodeGrpcServer.Serve(lsnDataNode); err != nil {
		log.Fatal("NameNode.grpcapi.service.go->NameNodeServeDataNode error:", err.Error())
	}
}
