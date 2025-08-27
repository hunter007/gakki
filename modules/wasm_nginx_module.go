package modules

import (
	"log/slog"
	"os/exec"
)

func innstallWASM(m *Module) error {
	cmd := exec.Command("./install-wasmtime.sh")
	cmd.Dir = m.Dir(m.version)
	output, err := cmd.CombinedOutput()
	slog.Info(string(output))
	if err != nil {
		return err
	}
	return nil
}

func setupWasmNginxModule() {
	module := &Module{
		name:        "wasm_nginx_module",
		tarFilename: "wasm-nginx-module",
		validVersions: []string{
			"0.7.0",
			"0.6.5",
			"0.6.4",
			"0.6.3",
			"0.6.2",
			"0.6.1",
			"0.6.0",
			"0.5.1",
			"0.5.0",
			"0.4.0",
			"0.3.0",
			"0.2.0",
			"0.1.0",
		},
		downloadTemplate: "https://github.com/api7/wasm-nginx-module/archive/refs/tags/%s.tar.gz",
		Install:          innstallWASM,
	}
	all[module.name] = module
}
