package search

// Package is the Go representation of a pkgparse package response
type Package struct {
	Description   string `json:"description"`
	Homepage      string `json:"homepage"`
	LatestVersion string `json:"latest_version"`
	License       string `json:"license"`
	Name          string `json:"name"`
	PackagePage   string `json:"package_page"`
	SourceRepo    bool   `json:"source_repo"`
	Tarball       string `json:"tarball"`
}

// Packages is a group of packages. Basically just for testing purposes.
type Packages []Package
