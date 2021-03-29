package blockmgr

import (
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

type blockMgr interface {
	LoadBlocks() (err error)
	RegisterDataNode() (result string, err error)
	Heartbeat() (result string, err error)
}

// BlockManager is a struct implementing interface "blockMgr"
type BlockManager struct {
	// Blocks contain some important information
	Blocks       []block
	BlocksOffset int

	// BlockFilesPathLocation is where Block files store
	BlockFilesPathLocation string
	BlockFilesPrefix       string

	// IDInSFS(int) => block idx of Blocks(block pool maintained by Mgr)
	IDMap map[int]int
}

// LoadBlocks is a method of interface blockMgr
func (blkmgr *BlockManager) LoadBlocks() (err error) {
	filesName, err := filepath.Glob(filepath.Join(blkmgr.BlockFilesPathLocation, "*"))
	if err != nil {
		log.Println("DataNode.BlockMgr.go->LoadBlocks error:", err.Error())
		return
	}

	for _, fileName := range filesName {
		if prefix, offset := blkmgr.BlockFilesPrefix, blkmgr.BlocksOffset; strings.HasPrefix(fileName, prefix) {
			// arr[0] is prefix, arr[1] is FilePathInSFS,
			// arr[2] is IDInSFS, arr[3] is OffsetInFile
			// arr[4] is LastBlockInFile
			arr := strings.Split(fileName, "_")
			FileNameBody := strings.TrimPrefix(fileName, prefix)
			IDInSFS, err := strconv.Atoi(arr[2])
			if err != nil {
				log.Println("DataNode.BlockMgr.go->LoadBlocks error:", err.Error())
				return err
			}
			OffsetInFile, err := strconv.Atoi(arr[3])
			if err != nil {
				log.Println("DataNode.BlockMgr.go->LoadBlocks error:", err.Error())
				return err
			}
			LastBlockInFile := (arr[4] == "1")

			// assignment about Blocks(block pool)
			blkmgr.Blocks[offset] = block{
				IDInSFS: uint32(IDInSFS), Status: stWELL,
				FileLocation:    blkmgr.BlockFilesPathLocation,
				FilePrefix:      blkmgr.BlockFilesPrefix,
				FileNameBody:    FileNameBody,
				FilePathInSFS:   arr[1],
				LastBlockInFile: LastBlockInFile,
				OffsetInFile:    uint32(OffsetInFile),
			}
			blkmgr.IDMap[IDInSFS] = offset
			blkmgr.BlocksOffset++
		}
	}
	return nil
}

// RegisterDataNode is a method of interface blockMgr (client)
func (blkmgr *BlockManager) RegisterDataNode() (result string, err error) {
	return
}

// Heartbeat is a method of interface blockMgr (client)
func (blkmgr *BlockManager) Heartbeat() (result string, err error) {
	return
}

// NewBlockManager is a constructor
// BlockFilesPathLocation default: "./"
// BlockFilesPrefix default: "data_"
func NewBlockManager(BlockFilesPathLocation, BlockFilesPrefix string) (blkmgr *BlockManager) {
	return &BlockManager{
		Blocks:                 make([]block, BLOCKSTOTALNUM),
		BlocksOffset:           0,
		BlockFilesPathLocation: BlockFilesPathLocation,
		BlockFilesPrefix:       BlockFilesPrefix,
		IDMap:                  make(map[int]int, BLOCKSTOTALNUM),
	}
}
