package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/skyleaworlder/Toy-HDFS/utils"
)

func Test_NewPacket(t *testing.T) {
	fd, _ := os.Open("D:\\CS·SB\\DS\\test.pdf")
	//fd, _ := os.Open("./Chunk_test.go")

	ctx := utils.NewPacketCtx(fd, 0, 1, true, 2)

	pkts := make([]utils.Packet, 3)

	ctx.Read(pkts)

	// test part
	for idx := range pkts {
		for jdx, val := range pkts[idx].PacketData {
			fmt.Println("jdx:", jdx, "; val:", val.Content)
		}
		fmt.Print("\n\n")
	}
}
