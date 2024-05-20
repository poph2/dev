package internal

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type BasePackage struct {
	Workspace      string
	Name           string
	CurrentVersion string
}

func (p BasePackage) SetupEnv() {

}

func (p BasePackage) Build() {

}

func (p BasePackage) Bump(releaseType ReleaseType) {

}

func (p BasePackage) Publish() {

}

func (p BasePackage) getCommitCount(tag string, subdir *string) int {

	command := "git rev-list --count " + tag + "..HEAD --pretty=oneline --count"

	if subdir != nil {
		command += " " + *subdir
	}

	out, err := RunCommand(command, p.Workspace)
	if err != nil {
		return -1
	}

	num, err := strconv.Atoi(strings.TrimSpace(out))

	if err != nil {
		return -1
	}

	return num
}

type NodeJsPackage struct {
	BasePackage
}

func (p NodeJsPackage) SetupEnv() {
}

func (p NodeJsPackage) Build() {

}

func (p NodeJsPackage) Bump(releaseType ReleaseType) {
	command := "npm version " + string(releaseType) + " --no-git-tag-version --no-commit-hooks --verbose`;"
	_, _ = RunCommand(command, p.Workspace)
}

func (p NodeJsPackage) Publish() {

}

type PythonPackage struct {
	BasePackage
}

func (p PythonPackage) SetupEnv() {
	// Check if ./venv exists
	_, err := os.Stat(filepath.Join(p.Workspace, "venv"))
	if err != nil {
		// Create a virtual environment
		_, _ = RunCommand("python3 -m venv venv", p.Workspace)
	}

	// install some tools
	_, _ = RunCommand("./venv/bin/pip3 install poetry poetry-bumpversion wheel twine", p.Workspace)
}

func (p PythonPackage) Build() {
	_, _ = RunCommand("./venv/bin/poetry build", p.Workspace)
}

func (p PythonPackage) Bump(releaseType ReleaseType) {
	_, _ = RunCommand("./venv/bin/poetry version "+string(releaseType), p.Workspace)
}

func (p PythonPackage) Publish() {

}
