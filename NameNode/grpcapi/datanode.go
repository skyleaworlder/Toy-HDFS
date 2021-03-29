package grpcapi

import (
	"context"
	"errors"
	"log"
	"net"
	"sync"

	"github.com/skyleaworlder/Toy-HDFS/NameNode/sfs"
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
	id := req.GetRegistration().GetID()
	if err := sfs.SFS.RegisterDataNode(id.GetHost(), net.ParseIP(id.GetIP()), id.GetTransferPort(), id.GetInfoPort()); err != nil {
		log.Println("NameNode.grpcapi.datanode.go->RegisterDataNode error:", err.Error())
		return nil, errors.New("NameNode.grpcapi.datanode.go->RegisterDataNode error:" + err.Error())
	}
	return &proto.RegisterDataNodeResponseProto{Registration: req.GetRegistration()}, nil
}

// SendHeartbeat is a method in interface "DataNodeNameNodeServer"
func (s *DataNodeNameNodeServer) SendHeartbeat(ctx context.Context, req *proto.HeartbeatRequestProto) (resp *proto.HeartbeatResponseProto, err error) {
	id, blocksInfo := req.GetDataNodeID(), req.GetBlocksInfo()
	sfs.SFS.ProcessHeartBeat(id, blocksInfo)
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
