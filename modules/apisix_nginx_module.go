package modules

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func setupApisixNginxModule() {
	module := &Module{
		name:        "apisix_nginx_module",
		tarFilename: "apisix-nginx-module",
		validVersions: []string{
			"1.19.2",
			"1.19.1",
			"1.19.0",
			"1.18.0",
			"1.17.0",
			"1.16.3",
			"1.16.2",
			"1.16.1",
			"1.16.0",
			"1.15.1",
			"1.15.0",
			"1.14.1",
			"1.14.0",
			"1.13.0",
			"1.12.0",
			"1.11.0",
			"1.10.0",
			"1.9.0",
			"1.8.0",
			"1.7.0",
			"1.6.0",
			"1.5.1",
			"1.5.0",
			"1.4.0",
			"1.3.1",
			"1.3.0",
			"1.2.0",
			"1.1.0",
			"1.0.0",
		},
		downloadTemplate: "https://github.com/api7/apisix-nginx-module/archive/refs/tags/%s.tar.gz",
		Patch:            patchForOpenresty,
	}
	all[module.name] = module
}

func patchForOpenresty(m *Module) error {
	cwd := m.Dir(m.version)
	openrestyDir := ""

	dirs, _ := os.ReadDir(filepath.Dir(cwd))
	for _, d := range dirs {
		if strings.Index(d.Name(), "openresty-") == 0 {
			openrestyDir = d.Name()
			break
		}
	}

	if openrestyDir == "" {
		err := fmt.Errorf("patch not found in %s", m.name)
		slog.Error(err.Error())
		return err
	}

	cmd := exec.Command("./patch.sh", "../../"+openrestyDir)
	cmd.Dir = fmt.Sprintf("%s%c%s", m.Dir(m.version), os.PathSeparator, "patch")
	output, err := cmd.CombinedOutput()
	if err != nil {
		slog.Info("patch error: " + string(output))
		return err
	}
	slog.Info("patch ok")
	return nil
}
