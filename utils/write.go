package utils

import (
	"io"
	"log"
	"net"
	"os"
	"runtime"
	"time"
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

func produceChunkByFd(fd *os.File, offsetInBlock int64, exit chan bool) {
	for offset, n, err := 0, 0, error(nil); err != io.EOF; offset += n {
		select {
		case <-exit:
			runtime.Goexit()

		default:
			// this is a bottle-neck
			log.Println("produceChunkByFd: get a new buf, offset is", offsetInBlock+int64(offset), len(hdfsOutPutBuffer.Buf))
			buf := make([]byte, CHUNKCONTENTSIZE)

			// if ReadAt cannot get a buffer filled with elem through fd
			// e.g. fd.ReadAt(buf, 12) --> [12 10 12 10 13 14 0 0 0 0 0]
			// then ReadAt will return err as io.EOF, instead of nil!!!
			//
			// ReadAt only return err as nil, when it can get a buffer like:
			// [12 10 12 10 13 14 67 69 77 100 2]
			//
			// so I still need to process a "defective" buffer.
			// I cannot do "
			// 	if err == io.EOF { break; }
			// " after the following assignment:
			n, err = fd.ReadAt(buf, offsetInBlock+int64(offset))
			chunk, _ := NewChunk(buf)

			//log.Print("begin lock.")
			hdfsOutPutBuffer.mtx.Lock()
			hdfsOutPutBuffer.Buf <- *chunk
			hdfsOutPutBuffer.BufferOffset++
			// self-listener
			if len(hdfsOutPutBuffer.Buf) == BUFCHUNKNUM {
				hdfsOutPutBuffer.done <- true
				hdfsOutPutBuffer.mtx.Unlock()
				time.Sleep(20 * time.Microsecond)
				break
			}
			hdfsOutPutBuffer.mtx.Unlock()
			//log.Println("end lock.")
		}
	}
}
