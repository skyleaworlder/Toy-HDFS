package sfs

import (
	"fmt"
	"net"
	"strconv"

	"github.com/skyleaworlder/Toy-HDFS/utils"
)

// InitFS is a function to init simple fs
func InitFS() {
	cfg := utils.ReadConfig("./", "slaves.yaml")
	// cfg.Get("slaves") => [map[]  map[]  map[] ... map[]]
	for _, datanode := range cfg.Get("slaves").([]map[string]string) {
		TransferPort, _ := strconv.Atoi(datanode["TransferPort"])
		InfoPort, _ := strconv.Atoi(datanode["InfoPort"])
		SFS.MachineArr[datanode["Host"]] = Machine{
			IP:           net.ParseIP(datanode["IP"]),
			TransferPort: uint32(TransferPort),
			InfoPort:     uint32(InfoPort),
		}
	}
	fmt.Println("SFS init process done.", SFS.MachineArr)
}
