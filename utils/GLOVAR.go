package utils

const (
	// BUFCHUNKNUM is num of chunk in buffer
	// Why it's 9? I think 126 % 9 = 0
	// (In fact, 126 % 7 also equal 0)
	// BUF write to Packet will process well
	BUFCHUNKNUM = 9
	// BUFSIZE is size of buffer used in write process
	BUFSIZE = CHUNKCONTENTSIZE * BUFCHUNKNUM

	// CHUNKCONTENTSIZE is size of chunk's content
	CHUNKCONTENTSIZE = 512
	// CHUNKCHECKSUMSIZE is size of chunk's checksum
	CHUNKCHECKSUMSIZE = 4
	// CHUNKTOTALSIZE equal size(checksum) + size(content)
	CHUNKTOTALSIZE = CHUNKCONTENTSIZE + CHUNKCHECKSUMSIZE

	// PACKETCHUNKNUM is num of chunk in packet
	PACKETCHUNKNUM = 126
	// PACKETCHUNKSIZE is increasing offset in file when writing & reading
	PACKETCHUNKSIZE = PACKETCHUNKNUM * CHUNKCONTENTSIZE
	// PACKETTOTALSIZE is size of packet
	PACKETTOTALSIZE = PACKETCHUNKNUM * CHUNKTOTALSIZE
)

var (
	hdfsOutPutBuffer       *Buffer      = newBuffer()
	hdfsPacketWaitingQueue *PacketQueue = NewPacketQueue()
	hdfsPacketACKQueue     *PacketQueue = NewPacketQueue()
)
