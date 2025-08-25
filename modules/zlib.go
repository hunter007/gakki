package modules

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"
)

func configZlib(m *Module) error {
	cmd := exec.Command("./configure")
	cmd.Dir = m.Dir(m.Version())
	out, err := cmd.Output()
	slog.Info(string(out))
	return err
}

func installZlib(m *Module) error {
	if err := configZlib(m); err != nil {
		return err
	}

	cmd := exec.Command("make")
	cmd.Dir = m.Dir(m.Version())
	out, err := cmd.Output()
	slog.Info(string(out))
	if err != nil {
		return err
	}

	f, err := os.Stat(m.Prefix())
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			if err = os.Mkdir(m.Prefix(), 0o755); err != nil {
				return err
			}
		}
		return err
	}
	if !f.IsDir() {
		return fmt.Errorf("%s is not directory", m.Prefix())
	}

	installCmd := exec.Command("make", "install", fmt.Sprintf("prefix=%s", m.Prefix()))
	installCmd.Dir = m.Dir(m.Version())
	installOut, err := installCmd.Output()
	slog.Info(string(installOut))

	return err
}

func setupZlib() {
	module := &Module{
		name:   "zlib",
		prefix: "/usr/local/zlib",
		validVersions: []string{
			"1.3.1",
			"1.3",
			"1.2.13",
			"1.2.12",
			"1.2.11",
			"1.2.10",
			"1.2.9",
			"1.2.8",
			"1.2.7.3",
			"1.2.7.2",
			"1.2.7.1",
			"1.2.7",
			"1.2.6.1",
			"1.2.6",
			"1.2.5.3",
			"1.2.5.2",
			"1.2.5.1",
			"1.2.5",
			"1.2.4.5",
			"1.2.4.4",
			"1.2.4.3",
			"1.2.4.2",
			"1.2.4.1",
			"1.2.4",
			"1.2.3.9",
			"1.2.3.8",
			"1.2.3.7",
			"1.2.3.6",
			"1.2.3.5",
			"1.2.3.4",
			"1.2.3.3",
			"1.2.3.2",
			"1.2.3.1",
			"1.2.3",
		},
		downloadTemplate: "https://www.zlib.net/fossils/zlib-%s.tar.gz",
		Install:          installZlib,
	}

	all[module.name] = module
}
