package modules

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"slices"
)

type (
	InstallModuleFunc = func(*Module) error
	Module            struct {
		name             string
		prefix           string
		version          string
		validVersions    []string
		downloadTemplate string
		dependences      []*Module
		Install          InstallModuleFunc
	}
)

func (m *Module) String() string {
	return m.name
}

func (m *Module) ListVersions() []string {
	return slices.Clone(m.validVersions)
}

func (m *Module) Download() error {
	if m.version == "" {
		return fmt.Errorf("no version")
	}

	fname := fmt.Sprintf("%s%c%s", dependentDir, os.PathSeparator, m.Filename(m.version))
	downloadUrl := fmt.Sprintf(m.downloadTemplate, m.version)
	cmd := exec.Command("wget", "--no-check-certificate", downloadUrl, "-O", fname)
	out, err := cmd.Output()
	slog.Info(string(out))
	return err
}

func (m *Module) Untar() error {
	if m.version == "" {
		return fmt.Errorf("no version")
	}

	fname := m.Filename(m.version)
	cmd := exec.Command("tar", "xzvf", fname)
	cmd.Dir = dependentDir
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func (m *Module) Filename(version string) string {
	return fmt.Sprintf("%s-%s.tar.gz", m.name, version)
}

func (m *Module) Dir(version string) string {
	return fmt.Sprintf("%s%c%s-%s", dependentDir, os.PathSeparator, m.name, version)
}

func (m *Module) VersionValid(version string) bool {
	return slices.Contains(m.validVersions, version)
}

func (m *Module) Prefix() string {
	return m.prefix
}

func (m *Module) Version() string {
	return m.version
}

func (m *Module) SetPrefix(prefix string) {
	m.prefix = prefix
}

func (m *Module) SetVersion(version string) error {
	if m.VersionValid(version) {
		m.version = version
		return nil
	}
	return fmt.Errorf("invalid version: %s", version)
}

func (m *Module) AddDependence(dep *Module) {
	m.dependences = append(m.dependences, dep)
}

func (m *Module) GetDependence(depName string) *Module {
	for _, dep := range m.dependences {
		if dep.name == depName {
			return dep
		}
	}
	return nil
}
