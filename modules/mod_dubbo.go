package modules

func setupModDubbo() {
	module := &Module{
		name: "mod_dubbo",
		validVersions: []string{
			"1.0.2",
			"1.0.1",
			"1.0.0",
		},
		downloadTemplate: "https://github.com/api7/mod_dubbo/archive/refs/tags/%s.tar.gz",
	}

	all[module.name] = module
}
