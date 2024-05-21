package cmd

import (
	"github.com/poph2/hive/internal"
	"github.com/spf13/cobra"
	"os"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the project",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := cmd.Flags().GetString("cwd")

		if cwd == "" {
			cwd, _ = os.Getwd()
		}

		project := internal.GetProject(cwd)
		project.Build()
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
