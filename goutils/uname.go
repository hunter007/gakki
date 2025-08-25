package goutils

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
)

type OS uint8

const (
	Unknow OS = 0
	Darwin OS = 1
	Linux  OS = 2
)

var curOS = Unknow

func init() {
	curOS = Uname()
}

func Uname() OS {
	cmd := exec.Command("uname")
	// if err := cmd.Run(); err != nil {
	// 	slog.Error(fmt.Sprintf("cannot get os: %s", err))
	// 	os.Exit(-1)
	// }
	out, err := cmd.Output()
	if err != nil {
		slog.Error(fmt.Sprintf("cannot get os: %s, output: %s", err, string(out)))
		os.Exit(-1)
	}

	osName := string(out)
	switch osName {
	case "Darwin":
		return Darwin
	case "Linux":
		return Linux
	default:
		return Unknow
	}
}
