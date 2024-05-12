package cmd

import (
	"github.com/poph2/hive/internal"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Start the project",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := cmd.Flags().GetString("cwd")
		_, _ = internal.RunCommand("npm i", cwd)
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}
