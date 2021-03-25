package utils

import "sync"

// Buffer is used in writing process
// BufferOffset represent size of content in buffer
type Buffer struct {
	Buf          []byte
	BufferOffset int
	mtx          *sync.Mutex
}

func newBuffer() *Buffer {
	return &Buffer{
		Buf:          make([]byte, BUFSIZE),
		BufferOffset: 0,
		mtx:          &sync.Mutex{},
	}
}
