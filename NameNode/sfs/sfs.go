package sfs

import (
	"errors"
	"log"
	"net"
)

// all files are stored in "/"
// there is no other directories.
var (
	// MAXFILENUM is max num of files in this simple fs
	MAXFILENUM = 10
	// MAXBLOCKSNUM is max num of blocks in this simple fs
	MAXBLOCKSNUM = 100
	// MAXMACHINENUM is max num of machines in this simple fs
	MAXMACHINENUM = 3

	// FileBlocksMap is a map(FileName => Blocks ID)
	FileBlocksMap = make(map[string][]int, MAXFILENUM)
	// BlockDataNodesMap is map(Block ID => DataNode IP)
	BlockDataNodesMap = make(map[int][]net.IP, MAXBLOCKSNUM)
	// BlockStatusArr is a []int
	BlockStatusArr = make([]int, MAXBLOCKSNUM)

	// MachineArr is a slice of available machine
	// and this fs only support 3-replicas cluster
	MachineArr = make(map[string]net.IP, MAXMACHINENUM)
)

// CreateFile is an api to create(register) file in fs
func CreateFile(FileName string, BlockNum uint32) (err error) {
	// check whether file exists or not
	if isExist := isFileExist(FileName); isExist {
		log.Println("NameNode.fs.fs.go->CreateFile error: file already exists")
		return errors.New("NameNode.fs.fs.go->CreateFile error: file already exists")
	}

	idxArr, err := getEmptyBlockIdx(BlockNum)
	if err != nil {
		log.Println("NameNode.fs.fs.go->CreateFile error:", err.Error())
		return errors.New("NameNode.fs.fs.go->CreateFile error: ")
	}

	FileNo := len(FileBlocksMap)
	for _, idx := range idxArr {
		// change BlockStatus Array
		BlockStatusArr[idx] = FileNo
	}
	// append record(Filename => Blocks ID)
	FileBlocksMap[FileName] = idxArr
	return nil
}

// GetFileBlocksLocation is an api to Get-Blocks-IP by FileName(FileName => Blocks ID => Blocks IP)
func GetFileBlocksLocation(FileName string) (FilePathInFS string, BlocksIdx []int, BlocksIP [][]net.IP, err error) {
	for file, blocks := range FileBlocksMap {
		if file == FileName {
			FilePathInFS = file
			copy(BlocksIdx, blocks)

			// copy each Block IPs by Block ID
			for _, BlockIdx := range BlocksIdx {
				var BlockIP []net.IP
				copy(BlockIP, BlockDataNodesMap[BlockIdx])
				BlocksIP = append(BlocksIP, BlockIP)
			}
			return
		}
	}
	return "", nil, nil, errors.New("NameNode.fs.fs.go->GetFileBlocksLocation error: cannot find file " + FileName)
}

// to get index of empty block
func getEmptyBlockIdx(BlockNum uint32) (res []int, err error) {
	res = []int{}
	var cnt uint32 = 0
	for idx, val := range BlockStatusArr {
		if cnt == BlockNum {
			return res, nil
		}
		if val == 0 {
			res = append(res, idx)
			cnt++
		}
	}
	return []int{}, errors.New("NameNode.fs.fs.go->getEmptyBlockIdx error: not enough space")
}

// to check whether file exists or not
func isFileExist(FileName string) (isExist bool) {
	_, isExist = FileBlocksMap[FileName]
	return
}
