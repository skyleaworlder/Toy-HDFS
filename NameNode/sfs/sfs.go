package sfs

import (
	"errors"
	"log"
	"net"
	"strings"

	"github.com/skyleaworlder/Toy-HDFS/proto"
)

const (
	// MAXFILENUM is max num of files in this simple fs
	MAXFILENUM = 10
	// MAXBLOCKSNUM is max num of blocks in this simple fs
	MAXBLOCKSNUM = 100
	// MAXMACHINENUM is max num of machines in this simple fs
	MAXMACHINENUM = 3
)

// Machine is a struct used in SFS
type Machine struct {
	IP           net.IP
	TransferPort uint32
	InfoPort     uint32
}

// SimpleFileSystem is a Primary File System
type SimpleFileSystem struct {
	// FileBlocksMap is a map(FileName => Blocks ID)
	// BlockDataNodesMap is map(Block ID => DataNode IP)
	// BlockStatusArr is a []int
	// MachineArr is a slice of available machine
	// and this fs only support 3-replicas cluster
	// hostname => net.IP & TransferPort & InfoPort
	FileBlocksMap     map[string][]int
	BlockDataNodesMap map[int][]net.IP
	BlockStatusArr    []int
	MachineArr        map[string]Machine
}

// all files are stored in "/"
// there is no other directories.
var (
	SFS *SimpleFileSystem = newSimpleFileSystem()
)

func newSimpleFileSystem() (SFS *SimpleFileSystem) {
	return &SimpleFileSystem{
		FileBlocksMap:     make(map[string][]int, MAXFILENUM),
		BlockDataNodesMap: make(map[int][]net.IP, MAXBLOCKSNUM),
		BlockStatusArr:    make([]int, MAXBLOCKSNUM),
		MachineArr:        make(map[string]Machine, MAXMACHINENUM),
	}
}

// CreateFile is an api to create(register) file in fs
func (SFS *SimpleFileSystem) CreateFile(FileName string, BlockNum uint32) (err error) {
	// check whether file exists or not
	if isExist := SFS.isFileExist(FileName); isExist {
		log.Println("NameNode.fs.fs.go->CreateFile error: file already exists")
		return errors.New("NameNode.fs.fs.go->CreateFile error: file already exists")
	}

	idxArr, err := SFS.getEmptyBlockIdx(BlockNum)
	if err != nil {
		log.Println("NameNode.fs.fs.go->CreateFile error:", err.Error())
		return errors.New("NameNode.fs.fs.go->CreateFile error: ")
	}

	FileNo := len(SFS.FileBlocksMap)
	for _, idx := range idxArr {
		// change BlockStatus Array
		SFS.BlockStatusArr[idx] = FileNo
	}
	// append record(Filename => Blocks ID)
	SFS.FileBlocksMap[FileName] = idxArr
	return nil
}

// GetFileBlocksLocation is an api to Get-Blocks-IP by FileName(FileName => Blocks ID => Blocks IP)
func (SFS *SimpleFileSystem) GetFileBlocksLocation(FileName string) (FilePathInFS string, BlocksIdx []int, BlocksIP [][]net.IP, err error) {
	for file, blocks := range SFS.FileBlocksMap {
		if file == FileName {
			FilePathInFS = file
			copy(BlocksIdx, blocks)

			// copy each Block IPs by Block ID
			for _, BlockIdx := range BlocksIdx {
				var BlockIP []net.IP
				copy(BlockIP, SFS.BlockDataNodesMap[BlockIdx])
				BlocksIP = append(BlocksIP, BlockIP)
			}
			return
		}
	}
	return "", nil, nil, errors.New("NameNode.fs.fs.go->GetFileBlocksLocation error: cannot find file " + FileName)
}

// RegisterDataNode is a function to register a datanode (server)
func (SFS *SimpleFileSystem) RegisterDataNode(Host string, IP net.IP, TransferPort, InfoPort uint32) (err error) {
	if err := checkGlovarSize(); err != nil {
		log.Println("NameNode.sfs.sfs.go->RegisterDataNode error:", err.Error())
		return err
	}
	SFS.MachineArr[Host] = Machine{IP: IP, TransferPort: TransferPort, InfoPort: InfoPort}
	return nil
}

// ProcessHeartBeat is a method to process Heartbeat signal (server)
func (SFS *SimpleFileSystem) ProcessHeartBeat(ID *proto.DataNodeIDProto, BlocksInfo []*proto.DataNodeBlockInfoProto) (err error) {
	if _, ok := SFS.MachineArr[ID.GetHost()]; !ok {
		return errors.New("NameNode.sfs.sfs.go->ProcessHeartBeat error: Host do not exist")
	}
	for _, blkInfo := range BlocksInfo {
		if err := processOneHeartBeat(blkInfo); err != nil {
			return err
		}
	}
	return nil
}

// to get index of empty block
func (SFS *SimpleFileSystem) getEmptyBlockIdx(BlockNum uint32) (res []int, err error) {
	res = []int{}
	var cnt uint32 = 0
	for idx, val := range SFS.BlockStatusArr {
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
func (SFS *SimpleFileSystem) isFileExist(FileName string) (isExist bool) {
	_, isExist = SFS.FileBlocksMap[FileName]
	return
}

func processOneHeartBeat(blkInfo *proto.DataNodeBlockInfoProto) (err error) {
	// block id do not exists in BlockDataNodesMap
	if _, ok := SFS.BlockDataNodesMap[int(blkInfo.GetBlockIdInSFS())]; !ok {
		return errors.New("NameNode.sfs.sfs.go->processOneHeartBeat error: Given BlockIdInSFS do not exist in BlockDataNodesMap")
	}

	FileNameInSFS := strings.Split(blkInfo.GetBlockNameBody(), "_")[0]
	if FileBlocks, ok := SFS.FileBlocksMap[FileNameInSFS]; !ok {
		// file do not exist in FileBlocksMap
		return errors.New("NameNode.sfs.sfs.go->processOneHeartBeat error: Given FileName do not in SFS.FileBlocksMap")
	} else if len(FileBlocks) <= int(blkInfo.GetBlockOffsetInFile()) {
		return errors.New("NameNode.sfs.sfs.go->processOneHeartBeat error: Given Block Offset In File is larger than Blocks Num")
	} else {
		// there would cause a lot of problems...
		return nil
	}
}

func checkGlovarSize() (err error) {
	if len(SFS.FileBlocksMap) >= MAXFILENUM {
		return errors.New("NameNode.sfs.sfs.go->checkGlovarSize error: FileBlocksMap outbound")
	} else if len(SFS.BlockDataNodesMap) >= MAXBLOCKSNUM {
		return errors.New("NameNode.sfs.sfs.go->checkGlovarSize error: BlockDataNodesMap outbound")
	} else if len(SFS.BlockStatusArr) >= MAXBLOCKSNUM {
		return errors.New("NameNode.sfs.sfs.go->checkGlovarSize error: BlockStatusArr outbound")
	} else if len(SFS.MachineArr) >= MAXMACHINENUM {
		return errors.New("NameNode.sfs.sfs.go->checkGlovarSize error: MachineArr outbound")
	} else {
		return nil
	}
}
