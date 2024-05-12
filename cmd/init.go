package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
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

func initializeNodeJsProject() {

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

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initializeNodeJsProject a new hive project",
	Run: func(cmd *cobra.Command, args []string) {
		initializeNodeJsProject()
		fmt.Println("init called...")
	},
}

var nodejsCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "initializeNodeJsProject a new hive project with Node.js",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

func init() {

	initCmd.AddCommand(nodejsCmd)

	rootCmd.AddCommand(initCmd)
}
