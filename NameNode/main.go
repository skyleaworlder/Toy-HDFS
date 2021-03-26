package main

import (
	"fmt"
	"log"
	"net"

	"github.com/skyleaworlder/Toy-HDFS/NameNode/grpcapi"
	"github.com/skyleaworlder/Toy-HDFS/utils"
)

func main() {
	cfg := utils.ReadConfig("./", "config-namenode.yaml")
	portRcvClient := cfg.GetString("namenode.port.rcv-Client")
	//portRcvDataNode := cfg.GetString("namenode.port.rcv-DataNode")

	ipSndClient := cfg.GetString("client.ip")
	portSndClient := cfg.GetString("client.port")
	//ipSndDataNode := cfg.GetString("datanode.ip")
	//portSndDataNode := cfg.GetString("datanode.port")

	// Parse TCP IP
	tcpIPRcvClient, err1 := net.ResolveTCPAddr("tcp4", portRcvClient)
	tcpIPSndClient, err2 := net.ResolveTCPAddr("tcp4", ipSndClient+portSndClient)
	if err1 != nil || err2 != nil {
		log.Fatal("ClientNode.main error: net.ResolveTCPAddr failed:", err1.Error(), err2.Error())
		return
	}

	// Generate listener
	lsnClient, err1 := net.ListenTCP("tcp4", tcpIPRcvClient)
	if err1 != nil {
		log.Fatal("ClientNode.main error: net.ResolveTCPAddr failed:", err1.Error())
		return
	}

	// Raise port listener service
	go grpcapi.NameNodeServeClient(lsnClient)

	// Raise NameNode "offsensive" service
	for {
		a := make(chan bool)
		<-a
		fmt.Println(tcpIPSndClient)
	}
}
