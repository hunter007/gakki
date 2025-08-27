package goutils

import (
	"os/exec"
	"runtime"
	"strconv"
)

func Nproc() uint8 {
	cmd := exec.Command("nproc")
	couNum, err := cmd.Output()
	if err != nil {
		return uint8(runtime.NumCPU())
	}
	cpuNum, _ := strconv.ParseInt(string(couNum), 10, 8)
	return uint8(cpuNum)
}
