package utils

import (
	pb "github.com/skyleaworlder/Toy-HDFS/proto"
)

// Chunk is a base data structure in hdfs
type Chunk struct {
	Checksum uint32
	Content  []byte
}

// ChunkPb is a data structure defined by protobuf
type ChunkPb struct {
	data *pb.ChunkProto
}

// NewChunk is a constructor of Chunk
func NewChunk(content []byte) (*Chunk, error) {
	chunk := new(Chunk)
	checksum, err := calcuChecksumByChunkContent(content)
	if err != nil {
		return &Chunk{}, err
	}
	chunk.Checksum = checksum
	chunk.Content = content
	return chunk, nil
}

// NewChunkPb is a constructor of Chunk defined by protobuf
func NewChunkPb(content []byte) (*ChunkPb, error) {
	data, _ := NewChunk(content)
	return &ChunkPb{
		data: &pb.ChunkProto{
			Checksum: data.Checksum,
			Content:  data.Content,
		},
	}, nil
}

// Chunk2ChunkProto is a translator
func Chunk2ChunkProto(chunk *Chunk) (chunkPb *pb.ChunkProto) {
	chunkPb.Checksum = chunk.Checksum
	chunkPb.Content = chunk.Content
	return
}
