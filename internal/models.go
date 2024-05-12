package internal

type ReleaseType string

const (
	Major ReleaseType = "major"
	Minor ReleaseType = "minor"
	Patch ReleaseType = "patch"
)

type PackageType string

const (
	Node   PackageType = "node"
	Python PackageType = "python"
)

type Package struct {
	Name string      `json:"name"`
	Type PackageType `json:"type"`
}

type Jack struct {
	Packages []Package `json:"packages"`
}

type PackageJSON struct {
	Jack Jack `json:"jack"`
}
