package modules

func setupWasmNginxModule() {
	module := &Module{
		name: "wasm-nginx-module",
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
	}
	all[module.name] = module
}
