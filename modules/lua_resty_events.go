package modules

func setupLuaRestyEvents() {
	module := &Module{
		name: "lua-resty-events",
		validVersions: []string{
			"0.3.1",
			"0.3.0",
			"0.2.1",
			"0.2.0",
			"0.1.6",
			"0.1.5",
			"0.1.4",
			"0.1.3",
			"0.1.2",
			"0.1.1",
			"0.1.0",
		},
		downloadTemplate: "https://github.com/Kong/lua-resty-events/archive/refs/tags/%s.tar.gz",
	}

	all[module.name] = module
}
