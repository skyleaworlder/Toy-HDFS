package utils

import (
	"hash/crc32"
)

func calcuChecksumByChunkContent(chunkContent []byte) (checksum uint32, err error) {
	// checksum's length is 4 bytes
	checksum = crc32.ChecksumIEEE(chunkContent)
	err = nil
	return
}
