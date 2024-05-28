package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type PackageJson struct {
	Name        string            `json:"name"`
	Version     string            `json:"version"`
	Description string            `json:"description"`
	Main        string            `json:"main"`
	Scripts     map[string]string `json:"scripts"`
	Keywords    []string          `json:"keywords"`
	Author      string            `json:"author"`
	License     string            `json:"license"`
}

type NodeJs struct {
	Project
}

func NewNodeJs(cwd string) NodeJs {
	return NodeJs{
		Project{
			Name:      "nodejs",
			Workspace: cwd,
			SetupEnvAction: Action{
				Run: []interface{}{"npm install"},
			},
			CleanAction: Action{
				Run: []interface{}{"rm -rf dist"},
			},
			BuildAction: Action{
				Run: []interface{}{"npm run build"},
			},
			BumpAction: Action{
				Run: []interface{}{
					func() bool {
						_, _ = RunCommand("npm version patch --no-git-tag-version --no-commit-hooks --verbose", cwd)
						return true
					},
					"npm version %s --no-git-tag-version --no-commit-hooks --verbose",
				},
			},
		},
	}
}

//func (p Nodejs) Build() {
//	_, _ = RunCommand("npm run build", p.Workspace)
//}
//
//func (p Nodejs) BumpRoot(cwd string, releaseType ReleaseType) {
//	command := "npm version " + string(releaseType) + " --no-git-tag-version --no-commit-hooks --verbose`;"
//	_, _ = RunCommand(command, cwd)
//}

//func (p Nodejs) Bump(releaseType ReleaseType) {
//
//	p.Build()
//
//	tag := p.getLatestTag()
//	commitCount := p.getCommitCount(tag, nil)
//	fmt.Println("...." + tag + "...")
//	fmt.Println("...." + strconv.Itoa(commitCount) + "...")
//
//	jsonFile, err := os.Open(filepath.Join(p.Workspace, "package.json"))
//
//	if err != nil {
//		fmt.Println(err)
//		panic(err)
//	}
//
//	defer func(jsonFile *os.File) {
//		err := jsonFile.Close()
//		if err != nil {
//			fmt.Println(err)
//		}
//	}(jsonFile)
//
//	byteValue, _ := io.ReadAll(jsonFile)
//
//	var packageJSON PackageJSON
//	err = json.Unmarshal(byteValue, &packageJSON)
//	if err != nil {
//		return
//	}
//
//	fmt.Println(packageJSON.Jack.Packages)
//
//	//var packagesToBump []Package
//
//	//for i, pkg := range packageJSON.Jack.Packages {
//	//	if pkg.Type == Node {
//	//
//	//		jsonData, _ := json.MarshalIndent(p, "", "  ")
//	//
//	//		workspace := "./packages/" + p.Name
//	//
//	//		commitCount := pkg.getCommitCount(tag, cwd, &workspace)
//	//
//	//		fmt.Println(i, string(jsonData), "....commit count: "+strconv.Itoa(commitCount)+"...")
//	//		if commitCount > 0 {
//	//			packagesToBump = append(packagesToBump, p)
//	//		}
//	//	}
//	//}
//	//
//	//command := "npm version " + string(releaseType) + " --no-git-tag-version --no-commit-hooks --verbose`;"
//	//_, _ = RunCommand(command, p.Workspace)
//}

//func BuildPackages(parentWorkingDir string, packages []Package) {
//	for _, p := range packages {
//		Build(filepath.Join(parentWorkingDir, "packages", p.Name), p.Type)
//	}
//}

func (p NodeJs) Init() {

	packageJson := PackageJson{
		Name:        "my-package",
		Version:     "1.0.0",
		Description: "",
		Main:        "index.js",
		Scripts:     map[string]string{"test": "echo \"Error: no test specified\""},
		Keywords:    []string{},
		Author:      "",
		License:     "ISC",
	}
	jsonData, err := json.MarshalIndent(packageJson, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	//content := `{
	//"name": "my-package",
	//"version": "1.0.0",
	//"description": "",
	//"main": "index.js",
	//"scripts": {
	//	"test": "echo \"Error: no test specified\" && exit 1"
	//},
	//"keywords": [],
	//"author": "",
	//"license": "ISC"
	//}`
	fmt.Println("Initializing a new Node.js project...")

	// Get WORKING_DIR env variable
	workingDir := os.Getenv("WORKING_DIR")

	if _, err := os.Stat(workingDir); os.IsNotExist(err) {
		errDir := os.MkdirAll(workingDir, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}

	// Create the package.json file
	file, err := os.Create(workingDir + "/package.json")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	// Write the content to the file
	err = os.WriteFile(workingDir+"/package.json", []byte(string(jsonData)), 0644)
	if err != nil {
		panic(err)
	}

}
