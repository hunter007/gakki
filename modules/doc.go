package modules

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

// dependentDir all modules will be downloaded to this directory
var (
	dependentDir string
)

var all = map[string]*Module{}

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		slog.Error("找不到用户目录，请创建并设置环境变量$HOME")
		os.Exit(-1)
	}

	dependentDir = fmt.Sprintf("%s%c%s", home, os.PathSeparator, ".deps")
	_, err = os.Stat(dependentDir)
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			if err := os.Mkdir(dependentDir, 0o755); err != nil {
				slog.Error(fmt.Sprintf("make dependent directory error: %s", err))
				os.Exit(-1)
			}
		} else {
			slog.Error(fmt.Sprintf("未知错误: %v, report it", err))
			os.Exit(-1)
		}
	}

	setupPcre2()
	setupZlib()
	setupEtcd()
	setupLimitTraffic()
	setupLuaRestyEvents()
	setupNgxMultiUpstreamModule()
	setupLuaVarNginxModule()
	setupModDubbo()
	setupOpenssl()
	setupWasmNginxModule()
	setupApisixNginxModule()
	setDeps()
	setupOpenresty()
}

func setDeps() {
}

func GetModule(name string) *Module {
	return all[name]
}
