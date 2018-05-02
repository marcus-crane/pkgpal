package search

import (
	"encoding/json"
	"sort"

	"github.com/marcus-crane/pkgpal/parsers"
)

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

// PypiPackage queries the /pypi/<package> endpoint for Python packages
func PypiPackage(name string) Package {
	response := Package{}

	body := Query("https://pkg.thingsima.de/pypi/" + name)

	unmarshallErr := json.Unmarshal(body, &response)
	if unmarshallErr != nil {
		panic(unmarshallErr)
	}
	return response
}

// ParseRequirements takes the file path for a requirements
// file, parses it, queries each package and returns
// a slice of Package structs
func ParseRequirements(path string) Packages {
	var packages Packages

	requirements := parsers.FeastRequirements(path)
	sort.Strings(requirements)

	for _, requirement := range requirements {
		packageResponse := PypiPackage(requirement)
		packages = append(packages, packageResponse)
	}

	return packages
}
