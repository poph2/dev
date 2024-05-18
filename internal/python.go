package internal

import (
	"os"
	"path/filepath"
)

type PythonP struct {
	Project
}

func (p PythonP) setupEnv() {
	// Check if ./venv exists
	_, err := os.Stat(filepath.Join(p.Workspace, "venv"))
	if err != nil {
		// Create a virtual environment
		_, _ = RunCommand("python3 -m venv venv", p.Workspace)
	}

	// install some tools
	_, _ = RunCommand("./venv/bin/pip3 install poetry poetry-bumpversion wheel twine", p.Workspace)
}

func (p PythonP) build() {
	//_, _ = RunCommand("./venv/bin/poetry build", p.workspace)
}

func (p PythonP) bump(releaseType ReleaseType) {

	_, _ = RunCommand("rm -rf *.egg-info dist build", p.Workspace)

	//currentVersion, _ := RunCommand("./venv/bin/poetry version -s", cwd)
}
