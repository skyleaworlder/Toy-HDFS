package sfs

import (
	"fmt"
	"net"

	"github.com/skyleaworlder/Toy-HDFS/utils"
)

// InitFS is a function to init simple fs
func InitFS() {
	cfg := utils.ReadConfig("./", "slaves.yaml")
	for k, v := range cfg.GetStringMap("slaves") {
		MachineArr[k] = net.ParseIP(v.(string))
	}
	fmt.Println("Hello World!", MachineArr)
}
