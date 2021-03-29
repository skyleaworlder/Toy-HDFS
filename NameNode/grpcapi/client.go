package grpcapi

import (
	"context"
	"log"
	"sync"

	"github.com/skyleaworlder/Toy-HDFS/NameNode/sfs"
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
	FilePathInFS, BlocksIdx, BlocksIP, err := sfs.SFS.GetFileBlocksLocation(req.GetFilePath())
	if err != nil {
		log.Println("NameNode.grpcapi.client.go->GetFileBlocksLocation error:", err.Error())
		return nil, err
	}

	BlocksNum := len(BlocksIdx)
	res := make([]*proto.BlockInfoProto, BlocksNum)
	for i := 0; i < BlocksNum; i++ {
		// process net.IP => string
		BlockIPs := make([]string, len(BlocksIP[i]))
		for idx, netIP := range BlocksIP[i] {
			BlockIPs[idx] = netIP.String()
		}

		res[i] = &proto.BlockInfoProto{
			FilePathInFS: FilePathInFS, BlockIndex: uint32(BlocksIdx[i]),
			LastBlockInFile: (i == BlocksNum), BlockIPs: BlockIPs,
		}
	}
	copy(resp.Blocks, res)
	return
}

// CreateFile is a method in interface "ClientNameNodeServer"
func (s *ClientNameNodeServer) CreateFile(ctx context.Context, req *proto.CreateFileRequestProto) (resp *proto.CreateFileResponseProto, err error) {
	err = sfs.SFS.CreateFile(req.GetFilePath(), req.GetBlockNum())
	if err != nil {
		log.Println("NameNode.grpcapi.client.go->CreateFile error:", err.Error())
		return
	}
	return &proto.CreateFileResponseProto{Result: "Create File " + req.GetFilePath() + " Success!"}, nil
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
