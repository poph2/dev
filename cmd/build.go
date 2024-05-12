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
		internal.Build(cwd, internal.Node)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
