package internal

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

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

type NodeJs struct {
	Project
}

type PythonP struct {
	Project
}

func NewNodeJs(cwd string) NodeJs {
	return NodeJs{
		Project{
			Name:      "nodejs",
			Workspace: cwd,
			SetupEnvAction: Action{
				Run: []interface{}{"npm install"},
			},
			CleanAction: Action{
				Run: []interface{}{"rm -rf dist"},
			},
			BuildAction: Action{
				Run: []interface{}{"npm run build"},
			},
			BumpAction: Action{
				Run: []interface{}{
					func() bool {
						_, _ = RunCommand("npm version patch --no-git-tag-version --no-commit-hooks --verbose", cwd)
						return true
					},
					"npm version %s --no-git-tag-version --no-commit-hooks --verbose",
				},
			},
		},
	}
}

func NewPythonP(cwd string) *PythonP {
	return &PythonP{
		Project{
			Name:      "python",
			Workspace: cwd,
			SetupEnvAction: Action{
				Check: func() bool {
					return dirExists(filepath.Join(cwd, "venv"))
				},
				Run: []interface{}{
					"python3 -m venv venv",
					"./venv/bin/pip3 install poetry poetry-bumpversion wheel twine",
					"./venv/bin/poetry install",
				},
			},
			CleanAction: Action{
				Run: []interface{}{"rm -rf dist"},
			},
			BuildAction: Action{
				Run: []interface{}{"./venv/bin/poetry build"},
			},
			BumpAction: Action{
				Run: []interface{}{"./venv/bin/poetry version %s"},
			},
		},
	}
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
