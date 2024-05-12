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

	if cwd != "" {
		cmd.Dir = cwd
	}

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
