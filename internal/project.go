package internal

import (
	"strconv"
	"strings"
)

type Projecter interface {
	SetupEnv()
	Clean()
	Build()
	Bump(releaseType ReleaseType)
	Publish()
	Init()
	getCommitCount(tag string, subdir *string) int
	getLatestTag() string
	gitCommit(message string)
	gitTag(tag string)
	gitPush(string)
}

type Project struct {
	Workspace      string
	Name           string
	CurrentVersion string
	CurrentTag     string

	SetupEnvAction Action
	CleanAction    Action
	BuildAction    Action
	BumpAction     Action
	PublishAction  Action
}

func (p Project) SetupEnv() {
	RunAction(p.SetupEnvAction, p.Workspace)
}

func (p Project) Clean() {
	RunAction(p.CleanAction, p.Workspace)
}

func (p Project) Build() {
	RunAction(p.BuildAction, p.Workspace)
}

func (p Project) Bump(releaseType ReleaseType) {
	RunAction(p.BumpAction, p.Workspace)
}

func (p Project) Publish() {
	RunAction(p.PublishAction, p.Workspace)
}

func (p Project) Init() {
	panic("implement me")
}

func (p Project) getCommitCount(tag string, subdir *string) int {

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

func (p Project) getLatestTag() string {
	out, err := RunCommand("git describe --tags --abbrev=0", p.Workspace)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(out)
}

func (p Project) gitCommit(message string) {
	_, _ = RunCommand("git commit -m '"+message+"' -a", p.Workspace)
}

func (p Project) gitTag(tag string) {
	_, _ = RunCommand("git tag -a "+tag+" -m '"+tag+"'", p.Workspace)
	_, _ = RunCommand("git push origin "+tag+" --no-verify", p.Workspace)
}

func (p Project) gitPush(string) {
	_, _ = RunCommand("git push --no-verify", p.Workspace)
}
