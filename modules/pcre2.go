package modules

func setupPcre2() {
	module := &Module{
		name: "pcre2",
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
	}

	all[module.name] = module
}
