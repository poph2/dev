package cmd

import (
	"fmt"
	"github.com/poph2/dev/internal/projects"
	"github.com/poph2/dev/internal/utilities"
	"github.com/spf13/cobra"
	"path/filepath"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new hive project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called...")
	},
}

var nodejsCmd = &cobra.Command{
	Use:   "nodejs [project-name]",
	Short: "Initialize a new hive project with Node.js",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		dir := utilities.GetCWD(cmd)

		cwd := filepath.Join(dir, projectName)

		fmt.Println("init called")
		project := projects.GetProject(projects.NewProjectOpts{Name: projectName, Workspace: cwd})
		project.Init()
		fmt.Println("init called")
	},
}

func init() {

	initCmd.AddCommand(nodejsCmd)

	rootCmd.AddCommand(initCmd)
}
