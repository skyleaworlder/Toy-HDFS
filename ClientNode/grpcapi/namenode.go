package grpcapi

import (
	"context"
	"fmt"
	"log"

	"github.com/skyleaworlder/Toy-HDFS/proto"
)

// PostCreateFileRequest is to send a request about creating a file in sfs
func PostCreateFileRequest(instrs []string, client proto.ClientNameNodeClient) {
	req := proto.CreateFileRequestProto{FilePath: instrs[1], BlockNum: 4}
	resp, err := client.CreateFile(context.Background(), &req)
	if err != nil {
		log.Println("ClientNode.grpcapi.namenode.go->PostCreateFileRequest error:", err.Error())
	}
	fmt.Println("Client>", resp)
}

// PostDeleteFileRequest is to send a request about deleting a file in sfs
func PostDeleteFileRequest(instrs []string, client proto.ClientNameNodeClient) {
	req := proto.DeleteFileRequestProto{FilePath: instrs[1]}
	resp, err := client.DeleteFile(context.Background(), &req)
	if err != nil {
		log.Println("ClientNode.grpcapi.namenode.go->PostDeleteFileRequest error:", err.Error())
	}
	fmt.Println("Client>", resp)
}
