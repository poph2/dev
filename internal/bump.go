package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

func bumpNodePackage(cwd string, releaseType ReleaseType) {
	command := "npm version " + string(releaseType) + " --no-git-tag-version --no-commit-hooks --verbose`;"
	_, _ = RunCommand(command, cwd)
}

func bump(cwd string, releaseType ReleaseType) {

	command := "npm version " + string(releaseType) + " --no-git-tag-version --no-commit-hooks --verbose`;"

	_, _ = RunCommand(command, cwd)
}

func bumpPackages(parentWorkingDir string, packages []Package, releaseType ReleaseType) {
	// build all the packages
	BuildPackages(parentWorkingDir, packages)

	// bump all the packages
	for _, p := range packages {
		bump(filepath.Join(parentWorkingDir, "packages", p.Name), releaseType)
	}

	// build the packages again
	BuildPackages(parentWorkingDir, packages)
}

func BumpProject(cwd string, releaseType ReleaseType) {

	// build the root package
	Build(cwd, Node)

	tag := GetLatestTag(cwd)
	commitCount := GetCommitCount(tag, cwd, nil)
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

	var packagesToBump []Package

	for i, p := range packageJSON.Jack.Packages {
		if p.Type == Node {

			jsonData, _ := json.MarshalIndent(p, "", "  ")

			workspace := "./packages/" + p.Name

			commitCount := GetCommitCount(tag, cwd, &workspace)

			fmt.Println(i, string(jsonData), "....commit count: "+strconv.Itoa(commitCount)+"...")
			if commitCount > 0 {
				packagesToBump = append(packagesToBump, p)
			}
		}
	}

	bumpPackages(cwd, packagesToBump, releaseType)

	// bump the root package
	bump(cwd, releaseType)

}
