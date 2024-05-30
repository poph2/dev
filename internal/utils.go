package internal

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
)

func dirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func GetCWD(cmd *cobra.Command) string {
	cwd, _ := cmd.Flags().GetString("cwd")

	if cwd == "" {
		cwd, _ = os.Getwd()
	}

	return cwd
}

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

func RunCommands(commands []string, cwd string) {
	for _, command := range commands {
		_, err := RunCommand(command, cwd)
		if err != nil {
			fmt.Println("Error running command: ", command)
			break
		}
	}
}

func GetProject(opt NewProjectOpts) Projecter {
	project := NewNodeJs(opt)
	fmt.Println("Project: ", project)
	return project
}
