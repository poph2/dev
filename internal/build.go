package internal

import "path/filepath"

type Command struct {
	Build string
	Bump  string
}

var nodejs = Command{
	Build: "npm run build",
	Bump:  "npm version patch",
}

func buildNodePackage(cwd string) {
	_, _ = RunCommand(nodejs.Build, cwd)
}

func buildPythonPackage(cwd string) {

}

func Build(cwd string, packageType PackageType) {

	switch packageType {
	case Node:
		buildNodePackage(cwd)
	case Python:
		buildPythonPackage(cwd)
	}
}

func BuildPackages(parentWorkingDir string, packages []Package) {
	for _, p := range packages {
		Build(filepath.Join(parentWorkingDir, "packages", p.Name), p.Type)
	}
}
