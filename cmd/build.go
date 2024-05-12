package cmd

import (
	"github.com/spf13/cobra"

	"github.com/poph2/hive/internal"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the project",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := cmd.Flags().GetString("cwd")
		_, _ = internal.RunCommand("nm run build", cwd)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
