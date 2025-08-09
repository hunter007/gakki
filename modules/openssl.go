package modules

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strconv"

	"github.com/hunter007/gakki/goutils"
)

func configOpenssl(m *Module) error {
	zlibModule := m.GetDependence("zlib")
	configArgsTemplate := `./config enable-fips shared zlib enable-camellia enable-seed enable-rfc3779 enable-cms enable-md2 enable-rc5 enable-weak-ssl-ciphers --prefix=%s --libdir=lib --with-zlib-lib=%s/lib --with-zlib-include=%s/include`
	configArgs := fmt.Sprintf(configArgsTemplate, m.Prefix(), zlibModule.Prefix(), zlibModule.Prefix())

	cmd := exec.Command("./config", configArgs)
	cmd.Dir = m.Dir(m.Version())
	cmd.Env = append(cmd.Environ(), fmt.Sprintf("LDFLAGS=-Wl,-rpath,%s/lib:%s/lib", zlibModule.Prefix(), m.Prefix()))

	out, err := cmd.Output()
	slog.Info(string(out))
	return err
}

func makeOpenssl(m *Module) error {
	zlibModule := m.GetDependence("zlib")

	cmd := exec.Command("make", "-j", strconv.Itoa(int(goutils.Nproc())), `LD_LIBRARY_PATH= CC="gcc"`)
	cmd.Dir = m.Dir(m.Version())
	cmd.Env = append(cmd.Environ(), fmt.Sprintf("LDFLAGS=-Wl,-rpath,%s/lib:%s/lib", zlibModule.Prefix(), m.Prefix()))

	out, err := cmd.Output()
	slog.Info(string(out))
	return err
}

func makeInstallOpenssl(m *Module) error {
	cmd := exec.Command("make", "install")
	cmd.Dir = m.Dir(m.Version())
	out, err := cmd.Output()
	slog.Info(string(out))
	return err
}

func copyOpensslConfFile(m *Module, confPath string) error {
	if _, err := os.Stat(confPath); err == nil {
		cmd := exec.Command("sudo", "cp", confPath, fmt.Sprintf("%s/ssl/openssl.cnf", m.Prefix()))
		out, err := cmd.Output()
		slog.Info(string(out))
		return err
	}
	return nil
}

func installFIPSForOpenssl(m *Module) error {
	opensslBin := fmt.Sprintf("%s/bin/openssl", m.Prefix())
	fipsModuleCnf := fmt.Sprintf("%s/ssl/fipsmodule.cnf", m.Prefix())
	modulePath := fmt.Sprintf("%s/lib/ossl-modules/fips.so", m.Prefix())

	cmd := exec.Command(opensslBin, "fipsinstall", "-out", fipsModuleCnf, "-module", modulePath)
	out, err := cmd.Output()
	slog.Info(string(out))
	return err
}

func modifyOpensslConf(m *Module) error {
	content := `'s@# .include fipsmodule.cnf@.include '"%s"'/ssl/fipsmodule.cnf@g; s/# \(fips = fips_sect\)/\1\nbase = base_sect\n\n[base_sect]\nactivate=1\n/g'`

	c := fmt.Sprintf(content, m.Prefix())
	cnf := fmt.Sprintf("%s/ssl/openssl.cnf", m.Prefix())

	cmd := exec.Command("sudo", "sed", "-i", c, cnf)
	out, err := cmd.Output()
	slog.Info(string(out))
	return err
}

func installOpenssl(m *Module) error {
	if err := configOpenssl(m); err != nil {
		return err
	}

	if err := makeOpenssl(m); err != nil {
		return err
	}

	if err := makeInstallOpenssl(m); err != nil {
		return err
	}
	// TODO:
	// if err := copyOpensslConfFile(m, confPath); err != nil {
	// 	return err
	// }

	if err := installFIPSForOpenssl(m); err != nil {
		return err
	}

	if err := modifyOpensslConf(m); err != nil {
		return err
	}

	return nil
}

func setupOpenssl() {
	module := &Module{
		name: "openssl",
		validVersions: []string{
			"3.5.2",
			"3.5.1",
			"3.5.0",
			"3.4.2",
			"3.4.1",
			"3.4.0",
			"3.3.4",
			"3.3.3",
			"3.3.2",
			"3.3.1",
			"3.3.0",
			"3.2.5",
			"3.2.4",
			"3.2.3",
			"3.2.2",
			"3.2.1",
			"3.2.0",
			"3.1.8",
			"3.1.7",
			"3.1.6",
			"3.1.5",
			"3.1.4",
			"3.1.3",
			"3.1.2",
			"3.1.1",
			"3.1.0",
			"3.0.17",
			"3.0.16",
			"3.0.15",
			"3.0.14",
			"3.0.13",
			"3.0.12",
			"3.0.11",
			"3.0.10",
			"3.0.9",
			"3.0.8",
			"3.0.7",
			"3.0.5",
			"3.0.4",
			"3.0.3",
			"3.0.2",
			"3.0.1",
			"3.0.0",
		},
		downloadTemplate: "https://github.com/openssl/openssl/releases/download/openssl-%s/openssl-%s.tar.gz",
		Install:          installOpenssl,
	}

	all[module.name] = module
	module.AddDependence(all["pcre2"])
	module.AddDependence(all["zlib"])
}
