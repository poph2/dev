package internal

import (
	"strconv"
	"strings"
)

//type PackageP interface {
//	SetupEnv()
//	Build()
//	Bump(releaseType ReleaseType)
//	Publish()
//}

type BasePackage struct {
	Workspace      string
	Name           string
	currentVersion string
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

}

func (p PythonPackage) Build() {

}

func (p PythonPackage) Bump(releaseType ReleaseType) {

}

func (p PythonPackage) Publish() {

}
