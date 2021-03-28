package main

import (
	"bufio"
	"log"
	"os"

	"github.com/skyleaworlder/Toy-HDFS/ClientNode/grpcapi"
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
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		instruction := input.Text()
		switch instrs := utils.ParseInstr(instruction); instrs[0] {
		case "create":
			grpcapi.PostCreateFileRequest(instrs, NameNodeClient)
		case "delete":
			grpcapi.PostDeleteFileRequest(instrs, NameNodeClient)
		}
	}

}
