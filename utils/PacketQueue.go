package utils

import (
	"errors"
	"sync"
)

// PacketQueue is a base data structure in hdfs
type PacketQueue struct {
	elements []Packet
	mtx      *sync.Mutex
}

// NewPacketQueue is a constructor
func NewPacketQueue() *PacketQueue {
	return &PacketQueue{
		elements: nil,
		mtx:      &sync.Mutex{},
	}
}

func (pq *PacketQueue) isEmpty() bool {
	return len(pq.elements) == 0
}

// GetLength is a method to return the num of elements
func (pq *PacketQueue) GetLength() int {
	return len(pq.elements)
}

// Push is to push a elem into queue
func (pq *PacketQueue) Push(elem Packet) error {
	pq.mtx.Lock()
	pq.elements = append(pq.elements, elem)
	pq.mtx.Unlock()
	return nil
}

// Pop is to pop a elem from queue
func (pq *PacketQueue) Pop() (out *Packet, err error) {
	defer pq.mtx.Unlock()
	pq.mtx.Lock()
	if pq.isEmpty() {
		return &Packet{}, errors.New("PacketQueue.Pop: Queue empty")
	}

	out = &pq.elements[0]
	pq.elements = pq.elements[1:]
	return out, nil
}
