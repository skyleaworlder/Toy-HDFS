package main

import (
	"io"
	"log"
	"os"
	"testing"

	"github.com/skyleaworlder/Toy-HDFS/utils"
)

func Test_NewChunk(t *testing.T) {
	// fd, err := os.Open("D:\\CS·SB\\DS\\test.pdf")
	fd, err := os.Open("../proto/chunk.pb.go")
	if err != nil {
		log.Println("test.Chunk_test.go->Test_NewChunk error:", err.Error())
		return
	}

	for offset, n := 0, 0; err != io.EOF; offset += n {
		buf := make([]byte, utils.CHUNKCONTENTSIZE)
		n, err = fd.ReadAt(buf, int64(offset))
		log.Println("hahaha:", buf, err)

		utils.NewChunk(buf)
		//log.Println("chunk:", chunk)
	}
}
