package utils

import (
	"errors"
	"io"
	"log"
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

// PacketReadWriter is an interface
type PacketReadWriter interface {
	Read(pktBuf []Packet)
	Write(pktBuf []Packet)
}

// PacketCtx is a user interface
// rw(io.ReadWriter) is usually file descriptor
// Buf's length equal CHUNKCONTENTSIZE
type PacketCtx struct {
	rw                io.ReadWriter
	OffsetInBlock     int64
	SeqNo             int32
	LastPacketInBlock bool
	DataLen           int32
	Buf               []byte
	EmptyBuf          []byte
}

// NewPacketCtx is a constructor
func NewPacketCtx(rw io.ReadWriter, OffsetInBlock int64, SeqNo int32, LastPacketInBlock bool, DataLen int32) (ctx *PacketCtx) {
	return &PacketCtx{
		rw: rw, OffsetInBlock: OffsetInBlock, SeqNo: SeqNo,
		LastPacketInBlock: LastPacketInBlock, DataLen: DataLen,
		Buf:      make([]byte, CHUNKCONTENTSIZE),
		EmptyBuf: make([]byte, CHUNKCONTENTSIZE),
	}
}

// n means success number
func (ctx *PacketCtx) Read(pkts []Packet) (n int, err error) {
	if len(pkts) == 0 {
		log.Println("utils.Packet.go->NewPacket.Read error: pkts empty buffer")
		return 0, errors.New("utils.Packet.go->NewPacket.Read error: pkts empty buffer")
	}

	// generate pkts, and put them into pkts([]Packet)
	for idx := range pkts {
		pkts[idx], err = getPacketContent(ctx)
		n++
	}
	return n, nil
}

// n means the number of write packets
func (ctx *PacketCtx) Write(pkts []Packet) (n int, err error) {
	if len(pkts) == 0 {
		log.Println("utils.Packet.go->NewPacket.Write error: pkts empty buffer")
		return 0, errors.New("utils.Packet.go->NewPacket.Write error: pkts empty buffer")
	}

	// push to waiting queue
	for _, pkt := range pkts {
		hdfsPacketWaitingQueue.Push(pkt)
		n++
	}
	return n, nil
}

func (ctx *PacketCtx) refreshBuf() {
	copy(ctx.Buf, ctx.EmptyBuf)
}

func getPacketContent(ctx *PacketCtx) (pkt Packet, err error) {
	pktp, _ := NewNullPacket(ctx.OffsetInBlock, ctx.SeqNo, ctx.LastPacketInBlock, ctx.DataLen)
	for len(pktp.PacketData) != PACKETCHUNKNUM {
		// fill ctx.Buf with content through fd
		// if ReadAt cannot get a buffer filled with elem through fd
		// e.g. fd.Read(buf, 12) --> [12 10 12 10 13 14 0 0 0 0 0]
		// then ReadAt will return err as nil
		ctx.refreshBuf()
		_, err := ctx.rw.Read(ctx.Buf)
		if err == io.EOF {
			pktp.mtx.Lock()
			pktp.PacketData = append(pktp.PacketData, Chunk{})
			pktp.PacketOffset++
			pktp.mtx.Unlock()
			continue
		} else if err == nil {
			// New a Chunk, put it into PacketData
			Chunk, err := NewChunk(ctx.Buf)
			//fmt.Println("Chunk:", Chunk, ctx.Buf)
			if err != nil {
				log.Println("utils.Packet.go->NewPacket.Read error: NewChunk error")
				return pkt, err
			}
			pktp.mtx.Lock()
			pktp.PacketData = append(pktp.PacketData, *Chunk)
			pktp.PacketOffset++
			pktp.mtx.Unlock()
		} else {
			log.Println("utils.Packet.go->getPacketContent error: ", err.Error())
			return *pktp, err
		}
	}
	return *pktp, nil
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

func newPacketHeader(offsetInBlock int64, seqNo int32, lastPacketInBlock bool, dataLen int32) *pb.PacketHeaderProto {
	return &pb.PacketHeaderProto{
		OffsetInBlock:     offsetInBlock,
		SeqNo:             seqNo,
		LastPacketInBlock: lastPacketInBlock,
		DataLen:           dataLen,
	}
}
