package internal

import (
	"os"
	"path/filepath"
)

type PythonProject struct {
	Project
}

func (p PythonProject) setupEnv() {
	// Check if ./venv exists
	_, err := os.Stat(filepath.Join(p.workspace, "venv"))
	if err != nil {
		// Create a virtual environment
		_, _ = RunCommand("python3 -m venv venv", p.workspace)
	}

	// install some tools
	_, _ = RunCommand("./venv/bin/pip3 install poetry poetry-bumpversion wheel twine", p.workspace)
}

func (p PythonProject) bump(cwd string, releaseType ReleaseType) {

	_, _ = RunCommand("rm -rf *.egg-info dist build", cwd)

	//currentVersion, _ := RunCommand("./venv/bin/poetry version -s", cwd)
}
