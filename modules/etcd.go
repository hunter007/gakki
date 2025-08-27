package modules

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
)

func setupEtcd() {
	module := &Module{
		name: "etcd",
		validVersions: []string{
			"3.6.4",
			"3.6.3",
			"3.6.2",
			"3.6.1",
			"3.6.0",
			"3.5.22",
			"3.5.21",
			"3.5.20",
			"3.5.19",
			"3.5.18",
			"3.5.17",
			"3.5.16",
			"3.5.15",
			"3.5.14",
			"3.5.13",
			"3.5.12",
			"3.4.37",
			"3.4.36",
			"3.4.35",
			"3.4.34",
			"3.4.33",
			"3.4.32",
			"3.4.31",
			"3.4.30",
		},
		downloadTemplate: "https://github.com/etcd-io/etcd/releases/download/v%s/etcd-v%s-%s-%s.zip",
		Install:          installEtcd,
	}

	all[module.name] = module
}

func installEtcd(m *Module) error {
	openrestyMod := GetModule("openresty")
	bin := fmt.Sprintf("%s%c%s", openrestyMod.Prefix(), os.PathSeparator, "bin")
	cmd := exec.Command("sudo", "cp", "./etcdctl", bin)
	cmd.Dir = m.Dir(m.version)

	output, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error(string(output))
		return err

	}
	slog.Info(string(output))
	return nil
}
