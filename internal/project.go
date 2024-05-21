package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type RootProject struct {
	Workspace      string
	Name           string
	CurrentVersion string
	CurrentTag     string
	Packages       []*RootProject
}

func (p RootProject) SetupEnv() {

}

func (p RootProject) clean() {

}

func (p RootProject) Build() {

}

func (p RootProject) Bump(releaseType ReleaseType) {

}

func (p RootProject) Publish() {

}

func (p RootProject) getCommitCount(tag string, subdir *string) int {

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

func (p RootProject) getLatestTag() string {
	out, err := RunCommand("git describe --tags --abbrev=0", p.Workspace)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(out)
}

func (p RootProject) gitCommit(message string) {
	_, _ = RunCommand("git commit -m '"+message+"' -a", p.Workspace)
}

func (p RootProject) gitTag(tag string) {
	_, _ = RunCommand("git tag -a "+tag+" -m '"+tag+"'", p.Workspace)
	_, _ = RunCommand("git push origin "+tag+" --no-verify", p.Workspace)
}

func (p RootProject) gitPush(string) {
	_, _ = RunCommand("git push --no-verify", p.Workspace)
}

//func (p RootProject) BuildProject() {
//	p.Build()
//
//	for _, pkg := range p.Packages {
//		pkg.Build()
//	}
//
//}

//func (p RootProject) BumpProject(releaseType ReleaseType) {
//
//	// tag := p.getLatestTag()
//	// commitCount := p.getCommitCount(tag, nil)
//
//	p.Bump(releaseType)
//
//	for _, pkg := range p.Packages {
//		pkg.Bump(releaseType)
//	}
//}

type NodeJs struct {
	RootProject
}

func (p NodeJs) Build() {
	fmt.Println("Building NodeJs package")
	_, _ = RunCommand("npm run build", p.Workspace)
}

func (p NodeJs) Bump(releaseType ReleaseType) {
	command := "npm version " + string(releaseType) + " --no-git-tag-version --no-commit-hooks --verbose`;"
	_, _ = RunCommand(command, p.Workspace)
}

func (p NodeJs) Publish() {

}

type PythonP struct {
	RootProject
}

func (p PythonP) SetupEnv() {
	// Check if ./venv exists
	_, err := os.Stat(filepath.Join(p.Workspace, "venv"))
	if err != nil {
		// Create a virtual environment
		_, _ = RunCommand("python3 -m venv venv", p.Workspace)
	}

	// install some tools
	_, _ = RunCommand("./venv/bin/pip3 install poetry poetry-bumpversion wheel twine", p.Workspace)
}

func (p PythonP) clean() {
	_, _ = RunCommand("rm -rf dist", p.Workspace)
}

func (p PythonP) Build() {
	_, _ = RunCommand("./venv/bin/poetry build", p.Workspace)
}

func (p PythonP) Bump(releaseType ReleaseType) {
	_, _ = RunCommand("./venv/bin/poetry version "+string(releaseType), p.Workspace)
}

func (p PythonP) Publish() {

}
