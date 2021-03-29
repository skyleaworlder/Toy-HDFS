package main

import (
	"fmt"
	"log"
	"net"

	"github.com/skyleaworlder/Toy-HDFS/NameNode/grpcapi"
	"github.com/skyleaworlder/Toy-HDFS/NameNode/sfs"
	"github.com/skyleaworlder/Toy-HDFS/utils"
)

func main() {
	// read config
	cfg := utils.ReadConfig("./", "config-namenode.yaml")
	sfs.InitFS()

	portRcvClient := cfg.GetString("namenode.port.rcv-Client")
	portRcvDataNode := cfg.GetString("namenode.port.rcv-DataNode")

	ipSndClient := cfg.GetString("client.ip")
	portSndClient := cfg.GetString("client.port")
	//ipSndDataNode := cfg.GetString("datanode.ip")
	//portSndDataNode := cfg.GetString("datanode.port")

	// Parse TCP IP
	tcpIPRcvClient, err1 := net.ResolveTCPAddr("tcp4", portRcvClient)
	tcpIPSndClient, err2 := net.ResolveTCPAddr("tcp4", ipSndClient+portSndClient)
	tcpIPRcvDataNode, err3 := net.ResolveTCPAddr("tcp4", portRcvDataNode)
	if err1 != nil || err2 != nil || err3 != nil {
		log.Fatal("ClientNode.main error: net.ResolveTCPAddr failed:", err1.Error(), err2.Error(), err3.Error())
		return
	}

	// Generate listener
	lsnClient, err1 := net.ListenTCP("tcp4", tcpIPRcvClient)
	lsnDataNode, err2 := net.ListenTCP("tcp4", tcpIPRcvDataNode)
	if err1 != nil || err2 != nil {
		log.Fatal("ClientNode.main error: net.ResolveTCPAddr failed:", err1.Error())
		return
	}

	// Raise port listener service
	go grpcapi.NameNodeServeClient(lsnClient)
	go grpcapi.NameNodeServeDataNode(lsnDataNode)

	// Raise NameNode "offsensive" service
	// now NameNode do not need this part
	for {
		a := make(chan bool)
		<-a
		fmt.Println(tcpIPSndClient)
	}
}
