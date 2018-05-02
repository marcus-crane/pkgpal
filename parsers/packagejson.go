package parsers

import (
	"encoding/json"
	"sort"
)

// PackageJSON defines the shape of a package json file.
// It ignores any keys we don't care about
type PackageJSON struct {
	Name            string     `json:"name"`
	Dependencies    Dependency `json:"dependencies"`
	DevDependencies Dependency `json:"devDependencies"`
}

// Dependency contains a bunch of key value pairs.
// We only care about the keys (package names)
type Dependency map[string]interface{}

// ParsePackageJSON parses a file into a PackageJSON struct
func ParsePackageJSON(path string) PackageJSON {
	file, _ := LoadFile(path, false)
	var pkgJSON PackageJSON
	err := json.Unmarshal(file, &pkgJSON)
	if err != nil {
		panic(err)
	}
	return pkgJSON
}

// SortDependencies takes a PackageJSON struct and returns
// the strings from the dependencies and devDependencies,
// sorted in alphabetical order.
func SortDependencies(pkgJSON PackageJSON) []string {
	dependencies := []string{}

	for key := range pkgJSON.Dependencies {
		dependencies = append(dependencies, key)
	}

	for key := range pkgJSON.DevDependencies {
		dependencies = append(dependencies, key)
	}

	sort.Strings(dependencies)
	return dependencies
}

// FeastPackageJSON takes in a file path for a package.json file
// and returns its dependencies (normal + dev) in alphabetical order
func FeastPackageJSON(path string) []string {
	pkgJSON := ParsePackageJSON(path)
	return SortDependencies(pkgJSON)
}
