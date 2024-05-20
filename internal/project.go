package internal

import (
	"strings"
)

type RootProject struct {
	BasePackage
	CurrentTag string
	Packages   []*BasePackage
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

func (p RootProject) BuildProject() {
	p.Build()

	for _, pkg := range p.Packages {
		pkg.Build()
	}

}

func (p RootProject) BumpProject(releaseType ReleaseType) {

	// tag := p.getLatestTag()
	// commitCount := p.getCommitCount(tag, nil)

	p.Bump(releaseType)

	for _, pkg := range p.Packages {
		pkg.Bump(releaseType)
	}
}

type NodeJsProject struct {
	RootProject
	NodeJsPackage
}

type PythonProject struct {
	RootProject
	PythonPackage
}
