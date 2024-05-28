package cmd

import (
	"fmt"
	"github.com/poph2/hive/internal"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initializeNodeJsProject a new hive project",
	Run: func(cmd *cobra.Command, args []string) {
		cwd := internal.GetCWD(cmd)
		project := internal.GetProject(cwd)
		project.Init()
		fmt.Println("init called...")
	},
}

var nodejsCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "initializeNodeJsProject a new hive project with Node.js",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

func init() {

	initCmd.AddCommand(nodejsCmd)

	rootCmd.AddCommand(initCmd)
}
