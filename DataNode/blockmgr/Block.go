package blockmgr

const (
	// BLOCKSTOTALNUM is the total number of blocks in one DataNode
	BLOCKSTOTALNUM = 10
)

// BLOCKSTATUS is enum type
type BLOCKSTATUS uint32

const (
	stIDLE    BLOCKSTATUS = 0
	stWELL    BLOCKSTATUS = 1
	stFAILED  BLOCKSTATUS = 2
	stREADING BLOCKSTATUS = 3
	stWRITING BLOCKSTATUS = 4
)

type block struct {
	IDInSFS uint32
	Status  BLOCKSTATUS

	// FileName = FileLocation + FilePrefix + FileNameBody
	FileLocation string
	FilePrefix   string
	FileNameBody string

	// FilePathInSFS is the path of file that consists of this block
	FilePathInSFS   string
	LastBlockInFile bool
	OffsetInFile    uint32
}
