package internal

type Lib struct {
	Workspace string
	Name      string
	IsRoot    bool `default:"false"`
}

func setDefault(lib *Lib) {
	if lib.IsRoot == false {
		lib.IsRoot = true
	}
}
