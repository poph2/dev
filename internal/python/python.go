package python

import (
	"os"
	"path/filepath"
)

func setupPythonEnv(cwd string) {
	// Check if ./venv exists
	_, err := os.Stat(filepath.Join(cwd, "venv"))
	if err != nil {
		// Create a virtual environment
		_, _ = RunCommand("python3 -m venv venv", cwd)
	}
}

func bumpPythonPackage(cwd string, releaseType ReleaseType) {

	// install some tools
	_, _ = RunCommand("./venv/bin/pip3 install poetry poetry-bumpversion wheel twine", cwd)
	_, _ = RunCommand("rm -rf *.egg-info dist build", cwd)

	currentVersion, _ := RunCommand("./venv/bin/poetry version -s", cwd)
}
