package internal

import (
	"fmt"
	"strconv"
	"strings"
)

func GetLatestTag(cwd string) string {
	fmt.Println("Getting latest tag")
	out, err := RunCommand("git describe --tags --abbrev=0", cwd)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(out)
}

func GetCommitCount(tag string, cwd string) int {
	out, err := RunCommand("git rev-list --count "+tag+"..HEAD --pretty=oneline --count", cwd)
	if err != nil {
		return -1
	}
	fmt.Println("Commit count: " + strings.TrimSpace(out))
	num, err := strconv.Atoi(strings.TrimSpace(out))
	fmt.Println(num)
	if err != nil {
		return -1
	}

	return num

}
