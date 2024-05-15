package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

type Nodejs struct {
	RootProject
}

func (p Nodejs) Build() {
	_, _ = RunCommand("npm run build", p.Workspace)
}

func (p Nodejs) Bump(releaseType ReleaseType) {

	p.Build()

	tag := p.getLatestTag()
	commitCount := p.getCommitCount(tag, nil)
	fmt.Println("...." + tag + "...")
	fmt.Println("...." + strconv.Itoa(commitCount) + "...")

	jsonFile, err := os.Open(filepath.Join(p.Workspace, "package.json"))

	if err != nil {
		fmt.Println(err)
		panic(err)
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

	for i, pkg := range packageJSON.Jack.Packages {
		if pkg.Type == Node {

			jsonData, _ := json.MarshalIndent(p, "", "  ")

			workspace := "./packages/" + p.Name

			commitCount := pkg.getCommitCount(tag, cwd, &workspace)

			fmt.Println(i, string(jsonData), "....commit count: "+strconv.Itoa(commitCount)+"...")
			if commitCount > 0 {
				packagesToBump = append(packagesToBump, p)
			}
		}
	}

	command := "npm version " + string(releaseType) + " --no-git-tag-version --no-commit-hooks --verbose`;"
	_, _ = RunCommand(command, p.Workspace)
}

//func BuildPackages(parentWorkingDir string, packages []Package) {
//	for _, p := range packages {
//		Build(filepath.Join(parentWorkingDir, "packages", p.Name), p.Type)
//	}
//}
