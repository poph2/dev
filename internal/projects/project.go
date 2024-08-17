package projects

import (
	"fmt"
	"github.com/poph2/dev/internal"
	"github.com/poph2/dev/internal/actions"
	"github.com/poph2/dev/internal/utilities"
	"strconv"
	"strings"
)

type Projecter interface {
	SetupEnv()
	Clean()
	Build()
	Bump(releaseType internal.ReleaseType)
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

	SetupEnvAction actions.Action
	CleanAction    actions.Action
	BuildAction    actions.Action
	BumpAction     actions.Action
	PublishAction  actions.Action
}

type NewProjectOpts struct {
	Name      string
	Workspace string
}

func (p Project) SetupEnv() {
	actions.RunAction(p.SetupEnvAction, p.Workspace)
}

func (p Project) Clean() {
	actions.RunAction(p.CleanAction, p.Workspace)
}

func (p Project) Build() {
	actions.RunAction(p.BuildAction, p.Workspace)
}

func (p Project) Bump(releaseType internal.ReleaseType) {
	actions.RunAction(p.BumpAction, p.Workspace)
}

func (p Project) Publish() {
	actions.RunAction(p.PublishAction, p.Workspace)
}

func (p Project) Init() {
	panic("implement me")
}

func (p Project) getCommitCount(tag string, subdir *string) int {

	command := "git rev-list --count " + tag + "..HEAD --pretty=oneline --count"

	if subdir != nil {
		command += " " + *subdir
	}

	out, err := utilities.RunCommand(command, p.Workspace)
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
	out, err := utilities.RunCommand("git describe --tags --abbrev=0", p.Workspace)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(out)
}

func (p Project) gitCommit(message string) {
	_, _ = utilities.RunCommand("git commit -m '"+message+"' -a", p.Workspace)
}

func (p Project) gitTag(tag string) {
	_, _ = utilities.RunCommand("git tag -a "+tag+" -m '"+tag+"'", p.Workspace)
	_, _ = utilities.RunCommand("git push origin "+tag+" --no-verify", p.Workspace)
}

func (p Project) gitPush(string) {
	_, _ = utilities.RunCommand("git push --no-verify", p.Workspace)
}

func GetProject(opt NewProjectOpts) Projecter {
	project := NewNodeJs(opt)
	fmt.Println("Project: ", project)
	return project
}
