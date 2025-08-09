package modules

import (
	"fmt"
	"log/slog"
	"os"
)

// dependentDir all modules will be downloaded to this directory
var (
	dependentDir string
	all          map[string]*Module
)

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		slog.Error("找不到用户目录，请创建并设置环境变量$HOME")
		os.Exit(-1)
	}

	dependentDir = fmt.Sprintf("%s%c%s", home, os.PathSeparator, ".deps")
	if err := os.Mkdir(dependentDir, 0o644); err != nil {
		slog.Error(fmt.Sprintf("未知错误: %s, report it", err))
		os.Exit(-1)
	}

	setupPcre2()
	setupZlib()
	setupApisixNginxModule()
	setupOpenssl()
	setDeps()
}

func setDeps() {
}

func GetModule(name string) *Module {
	return all[name]
}
