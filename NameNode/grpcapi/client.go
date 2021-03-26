package grpcapi

import (
	"context"
	"sync"

	"github.com/skyleaworlder/Toy-HDFS/proto"
)

// ClientNameNodeServer is a struct implement interface "ClientNameNodeServer"
type ClientNameNodeServer struct {
	Name     string
	Servefor string
	mtx      *sync.Mutex
}

// GetFileBlocksLocation is a method in interface "ClientNameNodeServer"
func (s *ClientNameNodeServer) GetFileBlocksLocation(ctx context.Context, req *proto.GetFileBlocksLocationRequestProto) (resp *proto.GetFileBlocksLocationResponseProto, err error) {
	return
}

// CreateFile is a method in interface "ClientNameNodeServer"
func (s *ClientNameNodeServer) CreateFile(ctx context.Context, req *proto.CreateFileRequestProto) (resp *proto.CreateFileResponseProto, err error) {
	return
}

// DeleteFile is a method in interface "ClientNameNodeServer"
func (s *ClientNameNodeServer) DeleteFile(ctx context.Context, req *proto.DeleteFileRequestProto) (resp *proto.DeleteFileResponseProto, err error) {
	return
}

// NewClientNameNodeServer is a constructor
func NewClientNameNodeServer() (s *ClientNameNodeServer) {
	return &ClientNameNodeServer{
		Name:     "NameNode",
		Servefor: "Client",
		mtx:      &sync.Mutex{},
	}
}
