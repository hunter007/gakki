package modules

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"

	"github.com/hunter007/gakki/goutils"
)

func setupLimitTraffic() {
	module := &Module{
		name:        "lua_resty_limit_traffic",
		tarFilename: "lua-resty-limit-traffic",
		validVersions: []string{
			"1.0.0",
			"0.08",
			"0.07",
		},
		downloadTemplate: "https://github.com/api7/lua-resty-limit-traffic/archive/refs/tags/v%s.tar.gz",
		Install:          installLimitTraffic,
	}

	all[module.name] = module
}

func installLimitTraffic(m *Module) error {
	openrestyMod := GetModule("openresty")
	openrestyDir := openrestyMod.Dir(openrestyMod.version)
	path := fmt.Sprintf("%s%c%s%c%s", openrestyDir, os.PathSeparator, "bundle", os.PathSeparator, "lua-resty-limit-traffic-0.09")

	if goutils.ExistDir(path) {
		if err := os.RemoveAll(path); err != nil {
			return err
		}
	}

	path2 := fmt.Sprintf("%s%c%s%c%s", openrestyDir, os.PathSeparator, "bundle", os.PathSeparator, "lua-resty-limit-traffic-1.0.0")
	if goutils.ExistDir(path2) {
		return nil
	}

	cmd := exec.Command("mv", m.Dir(m.version), fmt.Sprintf("%s%c%s", openrestyDir, os.PathSeparator, "bundle"))

	output, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error(string(output))
	}
	slog.Info(string(output))
	return nil
}
