package main

import (
	"context"
	"fmt"
	"log"

	"github.com/skyleaworlder/Toy-HDFS/proto"
	"github.com/skyleaworlder/Toy-HDFS/utils"
	"google.golang.org/grpc"
)

func main() {
	// Read config
	cfg := utils.ReadConfig("./", "config-client.yaml")

	//ipSndDataNode := cfg.GetString("datanode.ip")
	//portSndDataNode := cfg.GetString("datanode.port")
	ipSndNameNode := cfg.GetString("namenode.ip")
	portSndNameNode := cfg.GetString("namenode.port")

	// Parse TCP IP
	NameNodeConn, err1 := grpc.Dial(ipSndNameNode+portSndNameNode, grpc.WithInsecure())
	//DataNodeConn, err2 := grpc.Dial(ipSndDataNode+portSndDataNode, grpc.WithInsecure())
	if err1 != nil {
		log.Fatal("ClientNode.main error: net.ResolveTCPAddr failed:", err1.Error())
		return
	}

	// Generate Client
	NameNodeClient := proto.NewClientNameNodeClient(NameNodeConn)

	// Raise Client "offensive" service
	for {
		var instruction string
		fmt.Scan(&instruction)

		switch arr := utils.ParseInstr(instruction); arr[0] {
		case "create":
			req := proto.CreateFileRequestProto{FilePath: arr[1], BlockNum: 4}
			resp, err := NameNodeClient.CreateFile(context.Background(), &req)
			if err != nil {
				log.Println("ClientNode.main.go->main error:", err.Error())
			}
			fmt.Println(arr, resp)
		}
	}

}
