package utils

const (
	// BUFSIZE is size of buffer used in write process
	BUFSIZE = CHUNKCONTENTSIZE * 9

	// CHUNKCONTENTSIZE is size of chunk's content
	CHUNKCONTENTSIZE = 512
	// CHUNKCHECKSUMSIZE is size of chunk's checksum
	CHUNKCHECKSUMSIZE = 4
	// CHUNKTOTALSIZE equal size(checksum) + size(content)
	CHUNKTOTALSIZE = CHUNKCONTENTSIZE + CHUNKCHECKSUMSIZE

	// PACKETCHUNKNUM is num of chunk in packet
	PACKETCHUNKNUM = 126
	// PACKETTOTALSIZE is size of packet
	PACKETTOTALSIZE = PACKETCHUNKNUM * CHUNKTOTALSIZE
)

var (
	// Buffer is used in writing process
	Buffer []byte = make([]byte, BUFSIZE)
	// BufferOffset represent size of content in buffer
	BufferOffset int = 0
)
