package cmd

import (
	"github.com/poph2/hive/internal"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the project",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := cmd.Flags().GetString("cwd")

		project := internal.Nodejs{
			RootProject: internal.RootProject{Name: "nodejs", Workspace: cwd, IsRoot: true},
		}
		project.Build()
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
