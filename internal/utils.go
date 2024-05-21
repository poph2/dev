package internal

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func RunCommand(command string, cwd string) (string, error) {
	args := strings.Fields(command)
	cmd := exec.Command(args[0], args[1:]...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	fmt.Println("Running command: " + command)

	if cwd != "" {
		cmd.Dir = cwd
	}

	fmt.Println("Running command in directory: " + cmd.Dir)

	err := cmd.Run()

	var out string

	if err != nil {
		out = err.Error()
		fmt.Println(err)
	} else {
		out = stdout.String()
		fmt.Println(out)
	}
	return out, err
}

func GetProject(cwd string) NodeJs {
	project := NodeJs{
		RootProject{
			Name:           "nodejs",
			Workspace:      cwd,
			CurrentVersion: "1.0.0"},
	}
	fmt.Println("Project: ", project)
	return project
}
