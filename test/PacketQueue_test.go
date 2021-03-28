package main

import (
	"log"
	"testing"

	"github.com/skyleaworlder/Toy-HDFS/utils"
)

func Test_PacketQueue(t *testing.T) {
	pq := utils.NewPacketQueue()
	log.Println(pq.GetLength())
}
