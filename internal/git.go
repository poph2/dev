package internal

import (
	"strconv"
	"strings"
)

func GetLatestTag(cwd string) string {
	out, err := RunCommand("git describe --tags --abbrev=0", cwd)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(out)
}

func GetCommitCount(tag string, cwd string, workspace *string) int {

	command := "git rev-list --count " + tag + "..HEAD --pretty=oneline --count"

	if workspace != nil {
		command += " " + *workspace
	}

	out, err := RunCommand(command, cwd)
	if err != nil {
		return -1
	}

	num, err := strconv.Atoi(strings.TrimSpace(out))

	if err != nil {
		return -1
	}

	return num
}

func gitCommit(cwd string, message string) {
	_, _ = RunCommand("git commit -m '"+message+"' -a", cwd)
}

func gitTag(cwd string, tag string) {
	_, _ = RunCommand("git tag -a "+tag+" -m '"+tag+"'", cwd)
	_, _ = RunCommand("git push origin "+tag+" --no-verify", cwd)
}

func gitPush(cwd string) {
	_, _ = RunCommand("git push --no-verify", cwd)
}
