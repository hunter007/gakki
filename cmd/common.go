package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"slices"
	"strconv"
)

type OS uint8

const (
	Unknow OS = 0
	Darwin OS = 1
	Linux  OS = 2
)

var (
	curOS        = Unknow
	dependentDir string
)

func init() {
	curOS = Uname()
}

func Uname() OS {
	cmd := exec.Command("uname")
	out, err := cmd.Output()
	if err != nil {
		slog.Error(fmt.Sprintf("cannot get os: %s", err))
		os.Exit(-1)
	}

	osName := string(out)
	switch osName {
	case "Darwin":
		return Darwin
	case "Linux":
		return Linux
	default:
		return Unknow
	}
}

func Nproc() uint8 {
	cmd := exec.Command("nproc")
	if err := cmd.Run(); err != nil {
		slog.Error(fmt.Sprintf("cannot get CPU num: %s", err))
		os.Exit(-1)
	}
	couNum, err := cmd.Output()
	if err != nil {
		slog.Error(fmt.Sprintf("cannot get CPU num: %s", err))
		os.Exit(-1)
	}
	cpuNum, err := strconv.ParseInt(string(couNum), 10, 8)
	if err != nil {
		slog.Warn(fmt.Sprintf("cannot get CPU num: %s", err))
		return 1
	}
	return uint8(cpuNum)
}

func download(dep Dependent) error {
	fname := fmt.Sprintf("%s%c%s", dependentDir, os.PathSeparator, dep.FileName)
	cmd := exec.Command("wget", "--no-check-certificate", dep.Url, "-O", fname)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func Untar(path string) error {
	cmd := exec.Command("tar", "xzvf", path)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

type Dependent struct {
	Url      string
	FileName string
	Version  string
}

func (d Dependent) String() string {
	return fmt.Sprintf("%s%c%s", dependentDir, os.PathSeparator, d.FileName)
}

func (d Dependent) Dir() string {
	return fmt.Sprintf("%s%c%s", dependentDir, os.PathSeparator, d.FileName[:len(d.FileName)-7])
}

var nginxModule = Dependent{
	Url:      "",
	FileName: "",
	Version:  "1.19.2",
}

type Module struct {
	name             string
	validVersions    []string
	downloadTemplate string
	// Dependences []*Dependent
}

func (m *Module) String() string {
	return m.name
}

func (m *Module) download(version string) error {
	if !m.VersionValid(version) {
		return fmt.Errorf("invalid %s version: %s", m.name, version)
	}

	fname := fmt.Sprintf("%s%c%s", dependentDir, os.PathSeparator, m.filename(version))
	downloadUrl := fmt.Sprintf(m.downloadTemplate, version)
	cmd := exec.Command("wget", "--no-check-certificate", downloadUrl, "-O", fname)
	out, err := cmd.Output()
	slog.Info(string(out))
	return err
}

func (m *Module) untar(version string) error {
	if !m.VersionValid(version) {
		return fmt.Errorf("invalid %s version: %s", m.name, version)
	}

	fname := m.filename(version)
	cmd := exec.Command("tar", "xzvf", fname)
	cmd.Dir = dependentDir
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func (m *Module) filename(version string) string {
	return fmt.Sprintf("%s-%s.tar.gz", m.name, version)
}

func (m *Module) VersionValid(version string) bool {
	return slices.Contains(m.validVersions, version)
}
