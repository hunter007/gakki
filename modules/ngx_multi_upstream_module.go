package modules

func setupNgxMultiUpstreamModule() {
	module := &Module{
		name: "ngx_multi_upstream_module",
		validVersions: []string{
			"1.3.2",
			"1.3.1",
			"1.2.0",
			"1.1.1",
			"1.1.0",
			"1.0.1",
			"1.0.0",
		},
		downloadTemplate: "https://github.com/api7/ngx_multi_upstream_module/archive/refs/tags/%s.tar.gz",
		hasPatches:       true,
	}

	all[module.name] = module
}
