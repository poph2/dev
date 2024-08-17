package projects

import (
	"github.com/poph2/dev/internal/actions"
	"github.com/poph2/dev/internal/utilities"
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
			SetupEnvAction: actions.Action{
				Check: func() bool {
					return utilities.DirExists(filepath.Join(cwd, "venv"))
				},
				Run: []interface{}{
					"python3 -m venv venv",
					"./venv/bin/pip3 install poetry poetry-bumpversion wheel twine",
					"./venv/bin/poetry install",
				},
			},
			CleanAction: actions.Action{
				Run: []interface{}{"rm -rf dist"},
			},
			BuildAction: actions.Action{
				Run: []interface{}{"./venv/bin/poetry build"},
			},
			BumpAction: actions.Action{
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
		_, _ = utilities.RunCommand("python3 -m venv venv", p.Workspace)
	}

	// install some tools
	_, _ = utilities.RunCommand("./venv/bin/pip3 install poetry poetry-bumpversion wheel twine", p.Workspace)
}
