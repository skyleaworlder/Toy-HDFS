package main

import (
	"log"
	"testing"

	"github.com/skyleaworlder/Toy-HDFS/utils"
)

func Test_PacketQueue(t *testing.T) {
	pq := utils.NewPacketQueue()
	log.Println(pq.GetLength())
	pq.Push(utils.Packet{
		ChecksumField: []byte{1, 2, 3},
		ContentField:  []byte{4, 5, 6, 7, 8, 9, 10},
	})

	log.Println(pq.GetLength())
	elem, _ := pq.Pop()
	log.Println(elem)
	log.Println(pq.GetLength())
}
