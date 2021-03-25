package utils

import (
	"log"

	pb "github.com/skyleaworlder/Toy-HDFS/proto"
)

// Packet is base data structure in hdfs
type Packet struct {
	Header     *pb.PacketHeaderProto
	PacketData []Chunk
}

// PacketPb is a base data structure defined by protobuf
type PacketPb struct {
	data *pb.PacketProto
}

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

func newPacketHeader(offsetInBlock int64, seqNo int32, lastPacketInBlock bool, dataLen int32) *pb.PacketHeaderProto {
	return &pb.PacketHeaderProto{
		OffsetInBlock:     offsetInBlock,
		SeqNo:             seqNo,
		LastPacketInBlock: lastPacketInBlock,
		DataLen:           dataLen,
	}
}
