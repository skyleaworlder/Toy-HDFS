package utils

import (
	"log"
	"os"
	"sync"

	pb "github.com/skyleaworlder/Toy-HDFS/proto"
)

// Packet is base data structure in hdfs
// PacketData should be separate into checksumField and DataField,
// but I do not choose to implement that kind of packet format
//
// PacketOffset means the number of which Write function has appended,
// if PacketOffset equal PACKETCHUNKNUM(126), listen fiber will flush it.
type Packet struct {
	Header       *pb.PacketHeaderProto
	PacketData   []Chunk
	PacketOffset int
	mtx          *sync.Mutex
}

// PacketPb is a base data structure defined by protobuf
type PacketPb struct {
	data *pb.PacketProto
}

// NewNullPacket is a default construct
func NewNullPacket(offsetInBlock int64, seqNo int32, lastPacketInBlock bool, dataLen int32) (pkt *Packet, err error) {
	return &Packet{
		Header:       newPacketHeader(offsetInBlock, seqNo, lastPacketInBlock, dataLen),
		PacketData:   []Chunk{},
		PacketOffset: 0,
		mtx:          &sync.Mutex{},
	}, nil
}

// NewPacket is a constructor of Packet
// a fiber listen produceChunkByFd
//
// pkt = NewNullPacket() is better than pkt as global variable.
// global variable "pkt" without sync.Mutex is not thread-safe
// and do not support multi-thread generate pkt at the same time.
func NewPacket(fd *os.File, offsetInBlock int64, seqNo int32, lastPacketInBlock bool, dataLen int32) (pkt *Packet, newOffsetInBlock int64, err error) {
	exit := make(chan bool)
	defer func() { exit <- true }()
	go produceChunkByFd(fd, offsetInBlock, exit)

	pkt, _ = NewNullPacket(offsetInBlock, seqNo, lastPacketInBlock, dataLen)
	for {
		log.Println("NewPacket: in for")
		// buffer is filled with chunks
		// done(chan bool) is true, this fiber can get products from Buffer
		if <-hdfsOutPutBuffer.done {
			hdfsOutPutBuffer.mtx.Lock()
			for len(hdfsOutPutBuffer.Buf) > 0 {
				pkt.PacketData = append(pkt.PacketData, <-hdfsOutPutBuffer.Buf)
				pkt.PacketOffset++
			}
			hdfsOutPutBuffer.mtx.Unlock()
		}

		// packet is filled with chunks
		if pkt.PacketOffset == PACKETCHUNKNUM {
			newOffsetInBlock = offsetInBlock + int64(pkt.PacketOffset*CHUNKCONTENTSIZE)
			log.Println("return offset:", newOffsetInBlock)
			return pkt, newOffsetInBlock, nil
		}
	}
}

/*
// NewPacket is a constructor of Packet
func NewPacket(offsetInBlock int64, seqNo int32, lastPacketInBlock bool, dataLen int32, data []byte) (*Packet, error) {
	chunkArr := []Chunk{}
	content := make([]byte, CHUNKCONTENTSIZE)
	var i int = 0
	for ; i+CHUNKCONTENTSIZE < len(data); i += CHUNKCONTENTSIZE {
		byteNumToCopy := len(data) - i

		// if i+CHUNKCONTENTSIZE >= len(data), it means that there might be a number of bytes
		// (less than a chunk size) waiting to be put into a chunk.
		if 0 < byteNumToCopy && byteNumToCopy < CHUNKCONTENTSIZE {
			sliceCopy(data, i, content, 0, byteNumToCopy)
		} else {
			content = data[i : i+CHUNKCONTENTSIZE]
		}

		chunk, err := NewChunk(content)
		if err != nil {
			log.Println("utils.Packet.go->NewPacket error:", err.Error())
			return &Packet{}, err
		}
		chunkArr = append(chunkArr, *chunk)
	}

	return &Packet{
		Header:     newPacketHeader(offsetInBlock, seqNo, lastPacketInBlock, dataLen),
		PacketData: chunkArr,
	}, nil
}

// NewPacketPb is a constructor of PacketPb
// i fail to abstruct this two function...
func NewPacketPb(offsetInBlock int64, seqNo int32, lastPacketInBlock bool, dataLen int32, data []byte) (*PacketPb, error) {
	chunkProtoArr := []*pb.ChunkProto{}
	pkt, err := NewPacket(offsetInBlock, seqNo, lastPacketInBlock, dataLen, data)
	if err != nil {
		log.Println("utils.Packet.go->NewPacketPb error:", err.Error())
		return &PacketPb{}, nil
	}

	for _, chunk := range pkt.PacketData {
		chunkProtoArr = append(chunkProtoArr, Chunk2ChunkProto(&chunk))
	}
	return &PacketPb{
		data: &pb.PacketProto{
			Header:     newPacketHeader(offsetInBlock, seqNo, lastPacketInBlock, dataLen),
			PacketData: chunkProtoArr,
		},
	}, nil
}
*/

func newPacketHeader(offsetInBlock int64, seqNo int32, lastPacketInBlock bool, dataLen int32) *pb.PacketHeaderProto {
	return &pb.PacketHeaderProto{
		OffsetInBlock:     offsetInBlock,
		SeqNo:             seqNo,
		LastPacketInBlock: lastPacketInBlock,
		DataLen:           dataLen,
	}
}
