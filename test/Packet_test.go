package main

import (
	"log"
	"os"
	"testing"

	"github.com/skyleaworlder/Toy-HDFS/utils"
)

func Test_NewPacket(t *testing.T) {
	log.Println(utils.BUFCHUNKNUM)

	fd, _ := os.Open("D:\\CS·SB\\DS\\test.pdf")
	offset := 0
	for {
		log.Println("offset:", offset)
		_, newOffset, _ := utils.NewPacket(fd, int64(offset), 1, false, 5000)
		//log.Println("pkt:", pkt)
		log.Println("newOffset:", newOffset)
		offset = int(newOffset)
		log.Println("")
		//time.Sleep(1 * time.Second)
	}
}
