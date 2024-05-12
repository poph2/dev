package cmd

import (
	"github.com/poph2/hive/internal"
	"github.com/spf13/cobra"
)

var bumpCmd = &cobra.Command{
	Use:   "bump [major|minor|patch]",
	Short: "Bump the project version",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := cmd.Flags().GetString("cwd")
		internal.BumpProject(cwd, internal.ReleaseType(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(bumpCmd)
}
