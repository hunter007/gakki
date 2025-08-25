package modules

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"
)

func configPcre2(m *Module) error {
	cmd := exec.Command("./configure", fmt.Sprintf("--prefix=%s", m.Prefix()))
	cmd.Dir = m.Dir(m.Version())
	out, err := cmd.Output()
	slog.Info(string(out))
	return err
}

func installPcre2(m *Module) error {
	if err := configPcre2(m); err != nil {
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

	installCmd := exec.Command("make", "install")
	installCmd.Dir = m.Dir(m.Version())
	installOut, err := installCmd.Output()
	slog.Info(string(installOut))
	return nil
}

func setupPcre2() {
	module := &Module{
		name:   "pcre2",
		prefix: "/usr/local/pcre2",
		validVersions: []string{
			"10.45",
			"10.44",
			"10.43",
			"10.42",
			"10.41",
			"10.40",
			"10.39",
			"10.38",
			"10.37",
			"10.36",
			"10.35",
			"10.34",
			"10.33",
			"10.32",
			"10.31",
			"10.30",
			"10.23",
			"10.22",
			"10.21",
			"10.20",
			"10.10",
			"10.00",
		},
		downloadTemplate: "https://github.com/PCRE2Project/pcre2/releases/download/pcre2-%s/pcre2-%s.tar.gz",
		Install:          installPcre2,
	}

	all[module.name] = module
}

// CFLAGS='-O2 -Wall' ./configure --prefix=/opt/local
