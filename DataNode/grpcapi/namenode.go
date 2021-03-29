package grpcapi

import (
	"context"
	"fmt"
	"log"

	"github.com/skyleaworlder/Toy-HDFS/DataNode/blockmgr"
	"github.com/skyleaworlder/Toy-HDFS/proto"
)

// PostRegisterDataNode is to register
func PostRegisterDataNode(IP, Host string, TransferPort, InfoPort, BlockUsedNum uint32, client proto.DataNodeNameNodeClient) {
	id := &proto.DataNodeIDProto{IP: IP, Host: Host, TransferPort: TransferPort, InfoPort: InfoPort}
	storage := &proto.StorageReportProto{BlocksTotalNum: blockmgr.BLOCKSTOTALNUM, BlocksUsedNum: BlockUsedNum}
	datanodeRegistration := &proto.DataNodeRegistrationProto{ID: id, Storage: storage}
	req := proto.RegisterDataNodeRequestProto{Registration: datanodeRegistration}
	resp, err := client.RegisterDataNode(context.Background(), &req)
	if err != nil {
		log.Println("DataNode.grpcapi.namenode.go->PostRegisterDataNode error:", err.Error())
		return
	}
	fmt.Println("DataNode>", resp)
}

// PostHeartBeat is to send heartbeat
func PostHeartBeat(IP, Host string, TransferPort, InfoPort uint32, blkmgr *blockmgr.BlockManager, client proto.DataNodeNameNodeClient) {
	id := &proto.DataNodeIDProto{IP: IP, Host: Host, TransferPort: TransferPort, InfoPort: InfoPort}
	datanodeBlocksInfo := []*proto.DataNodeBlockInfoProto{}
	for _, blk := range blkmgr.Blocks {
		datanodeBlocksInfo = append(datanodeBlocksInfo, &proto.DataNodeBlockInfoProto{
			BlockIdInSFS:    blk.IDInSFS,
			BlockNamePrefix: blk.FilePrefix,
			BlockNameBody:   blk.FileNameBody,
			BlockFailed:     false,
		})
	}
	req := proto.HeartbeatRequestProto{DataNodeID: id, BlocksInfo: datanodeBlocksInfo}
	resp, err := client.SendHeartbeat(context.Background(), &req)
	if err != nil {
		log.Println("DataNode.grpcapi.namenode.go->PostHeartBeat error:", err.Error())
		return
	}
	fmt.Println("DataNode>", resp)
}
