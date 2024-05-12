package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/poph2/hive/internal"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

type ReleaseType string

const (
	Major ReleaseType = "major"
	Minor ReleaseType = "minor"
	Patch ReleaseType = "patch"
)

type Package struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Jack struct {
	Packages []Package `json:"packages"`
}

type PackageJSON struct {
	Jack Jack `json:"jack"`
}

func bump(cwd string, releaseType ReleaseType) {
	tag := internal.GetLatestTag(cwd)
	commitCount := internal.GetCommitCount(tag, cwd)
	fmt.Println("...." + tag + "...")
	fmt.Println("...." + strconv.Itoa(commitCount) + "...")

	jsonFile, err := os.Open(filepath.Join(cwd, "package.json"))

	if err != nil {
		fmt.Println(err)
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(jsonFile)

	byteValue, _ := io.ReadAll(jsonFile)

	var packageJSON PackageJSON
	err = json.Unmarshal(byteValue, &packageJSON)
	if err != nil {
		return
	}

	fmt.Println(packageJSON.Jack.Packages)

}

var bumpCmd = &cobra.Command{
	Use:   "bump [major|minor|patch]",
	Short: "Bump the project version",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := cmd.Flags().GetString("cwd")
		bump(cwd, ReleaseType(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(bumpCmd)
}
