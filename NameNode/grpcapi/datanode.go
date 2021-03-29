package grpcapi

import (
	"context"
	"sync"

	"github.com/skyleaworlder/Toy-HDFS/proto"
)

// DataNodeNameNodeServer is a struct implementing interface "DataNodeNameNodeServer"
type DataNodeNameNodeServer struct {
	Name     string
	Servefor string
	mtx      *sync.Mutex
}

// RegisterDataNode is a method in interface "DataNodeNameNodeServer"
func (s *DataNodeNameNodeServer) RegisterDataNode(ctx context.Context, req *proto.RegisterDataNodeRequestProto) (resp *proto.RegisterDataNodeResponseProto, err error) {
	return
}

// SendHeartbeat is a method in interface "DataNodeNameNodeServer"
func (s *DataNodeNameNodeServer) SendHeartbeat(ctx context.Context, req *proto.HeartbeatRequestProto) (resp *proto.HeartbeatResponseProto, err error) {
	return
}

// NewDataNodeNameNodeServer is a constructor
func NewDataNodeNameNodeServer() (s *DataNodeNameNodeServer) {
	return &DataNodeNameNodeServer{
		Name:     "NameNode",
		Servefor: "DataNode",
		mtx:      &sync.Mutex{},
	}
}
