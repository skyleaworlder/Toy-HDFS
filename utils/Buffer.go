package utils

import "sync"

// Buffer is used in writing process
// BufferOffset represent the number of elem in buffer
type Buffer struct {
	Buf          chan Chunk
	BufferOffset int
	mtx          *sync.Mutex
}

func newBuffer() *Buffer {
	return &Buffer{
		Buf:          make(chan Chunk, BUFCHUNKNUM),
		BufferOffset: 0,
		mtx:          &sync.Mutex{},
	}
}
