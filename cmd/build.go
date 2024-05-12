package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

func runCommand(command string, cwd string) (string, error) {
	args := strings.Fields(command)
	cmd := exec.Command(args[0], args[1:]...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if cwd != "" {
		cmd.Dir = cwd
	}

	err := cmd.Run()
	if err != nil {
		return stderr.String(), err
	}
	return out.String(), nil
}

func BuildProject(cwd string) {
	output, err := runCommand("npm run build", cwd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(output)
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the project",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := cmd.Flags().GetString("cwd")
		BuildProject(cwd)
	},
}

func init() {
	buildCmd.Flags().StringP("cwd", "c", "", "The current working directory")
	rootCmd.AddCommand(buildCmd)
}
