package goutils

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func Nproc() uint8 {
	cmd := exec.Command("nproc")
	if err := cmd.Run(); err != nil {
		slog.Error(fmt.Sprintf("cannot get CPU num: %s", err))
		os.Exit(-1)
	}
	couNum, err := cmd.Output()
	if err != nil {
		slog.Error(fmt.Sprintf("cannot get CPU num: %s", err))
		os.Exit(-1)
	}
	cpuNum, err := strconv.ParseInt(string(couNum), 10, 8)
	if err != nil {
		slog.Warn(fmt.Sprintf("cannot get CPU num: %s", err))
		return uint8(runtime.NumCPU())
	}

	return uint8(cpuNum)
}
