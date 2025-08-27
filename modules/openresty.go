package modules

func installOpenresty(m *Module) error {
	return nil
}

func setupOpenresty() {
	module := &Module{
		name:   "openresty",
		prefix: "/usr/local/openresty",
		validVersions: []string{
			"1.27.1.2",
			"1.25.3.2",
			"1.25.3.1",
			"1.21.4.4",
			"1.21.4.3",
			"1.21.4.2",
			"1.21.4.1",
			"1.19.9.2",
			"1.19.9.1",
			"1.19.3.2",
			"1.19.3.1",
		},
		downloadTemplate: "https://openresty.org/download/openresty-%s.tar.gz",
		Install:          installOpenresty,
	}

	all[module.name] = module
	module.AddDependence(all["pcre2"])
	module.AddDependence(all["zlib"])
}
