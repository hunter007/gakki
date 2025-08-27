package modules

func setupLimitTraffic() {
	module := &Module{
		name:        "lua_resty_limit_traffic",
		tarFilename: "lua-resty-limit-traffic",
		validVersions: []string{
			"1.0.0",
			"0.08",
			"0.07",
		},
		downloadTemplate: "https: //github.com/api7/lua-resty-limit-traffic/archive/refs/tags/v%s.tar.gz",
	}

	all[module.name] = module
}
