package cmd

import (
	"fmt"
	"github.com/poph2/hive/internal"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean the project",
	Run: func(cmd *cobra.Command, args []string) {
		cwd := internal.GetCWD(cmd)

		project := internal.GetProject(internal.NewProjectOpts{Workspace: cwd})
		project.Clean()
	},
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the project",
	Run: func(cmd *cobra.Command, args []string) {
		cwd := internal.GetCWD(cmd)

		project := internal.GetProject(internal.NewProjectOpts{Workspace: cwd})

		fmt.Println("Building project...")

		project.Clean()
		project.Build()
	},
}

var bumpCmd = &cobra.Command{
	Use:   "bump [major|minor|patch]",
	Short: "Bump the project version",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cwd := internal.GetCWD(cmd)
		releaseType := internal.ReleaseType(args[0])

		project := internal.GetProject(internal.NewProjectOpts{Workspace: cwd})

		project.Clean()
		project.Build()
		project.Bump(releaseType)
	},
}

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Start the project",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := cmd.Flags().GetString("cwd")
		_, _ = internal.RunCommand("npm i", cwd)
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(bumpCmd)
	rootCmd.AddCommand(upCmd)
}
