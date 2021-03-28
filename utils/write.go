package utils

import (
	"log"
	"net"
	"os"
)

func sliceCopy(src []byte, srcOffset int, dst []byte, dstOffset, copyLen int) {
	for i := 0; i < copyLen; i++ {
		dst[dstOffset+i] = src[srcOffset+i]
	}
}

// HDFSOutputStream can write
type HDFSOutputStream struct {
	Target                 net.TCPAddr
	writtenPacketTotalSize int
	writtenPacketDataSize  int
	lastPacketSize         int
	pipelineBreak          bool
	done                   bool
}

// Write is user interface
// a function, use a stream-object to upload file to hdfs
// 1. communicate with namenode
//		build a conn between name node and client,
//		send localPath, query blocks and get metadata
// 2. build a conn between data node and client
// 3. for {
//		get NewPacket as "pkt" from file descriptor.
//		conn.Write(pkt)
// }
func (ops *HDFSOutputStream) Write(localPath string) {
	fd, err := os.Open(localPath)
	if err != nil {
		log.Println("utils.write.go->Test_NewChunk error: open local file failed")
		log.Println("utils.write.go->Test_NewChunk error:", err.Error())
		return
	}
	log.Println(fd)

}
