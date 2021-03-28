package utils

import (
	"fmt"
	"strings"
)

// ParseInstr is a function to process instrucion
func ParseInstr(instr string) (instrs []string) {
	instrs = strings.Split(instr, ",")
	fmt.Println(instrs, instr)
	return instrs
}
