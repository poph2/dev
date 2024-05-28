package internal

import (
	"os"
	"path/filepath"
)

type PythonP struct {
	Project
}

func NewPythonP(cwd string) *PythonP {
	return &PythonP{
		Project{
			Name:      "python",
			Workspace: cwd,
			SetupEnvAction: Action{
				Check: func() bool {
					return dirExists(filepath.Join(cwd, "venv"))
				},
				Run: []interface{}{
					"python3 -m venv venv",
					"./venv/bin/pip3 install poetry poetry-bumpversion wheel twine",
					"./venv/bin/poetry install",
				},
			},
			CleanAction: Action{
				Run: []interface{}{"rm -rf dist"},
			},
			BuildAction: Action{
				Run: []interface{}{"./venv/bin/poetry build"},
			},
			BumpAction: Action{
				Run: []interface{}{"./venv/bin/poetry version %s"},
			},
		},
	}
}

func (p PythonP) SetupEnv() {
	// Check if ./venv exists
	_, err := os.Stat(filepath.Join(p.Workspace, "venv"))
	if err != nil {
		// Create a virtual environment
		_, _ = RunCommand("python3 -m venv venv", p.Workspace)
	}

	// install some tools
	_, _ = RunCommand("./venv/bin/pip3 install poetry poetry-bumpversion wheel twine", p.Workspace)
}
