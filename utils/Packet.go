package utils

import (
	pb "github.com/skyleaworlder/Toy-HDFS/proto"
)

// Packet is a base data structure in hdfs
type Packet struct {
	Header        pb.PacketHeaderProto
	ChecksumField []byte
	ContentField  []byte
}

// NewPacket is a constructor
func NewPacket(offsetInBlock int64, seqNo int32, lastPacketInBlock bool, dataLen int32, checksumField, contentField []byte) (pkt *Packet) {
	return &Packet{
		Header:        *newPacketHeader(offsetInBlock, seqNo, lastPacketInBlock, dataLen),
		ChecksumField: checksumField,
		ContentField:  contentField,
	}
}

func newPacketHeader(offsetInBlock int64, seqNo int32, lastPacketInBlock bool, dataLen int32) *pb.PacketHeaderProto {
	return &pb.PacketHeaderProto{
		OffsetInBlock:     offsetInBlock,
		SeqNo:             seqNo,
		LastPacketInBlock: lastPacketInBlock,
		DataLen:           dataLen,
	}
}
