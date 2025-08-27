package modules

func setupLuaVarNginxModule() {
	module := &Module{
		name: "lua_var_nginx_module",
		validVersions: []string{
			"0.5.3",
			"0.5.2",
			"0.5.1",
		},
		downloadTemplate: "https://github.com/api7/lua-var-nginx-module/archive/refs/tags/v%s.tar.gz",
	}
	all[module.name] = module
}
