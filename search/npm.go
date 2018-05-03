package search

import (
	"encoding/json"
	"sort"

	"github.com/marcus-crane/pkgpal/parsers"
)

// PypiPackage queries the /pypi/<package> endpoint for Python packages
func NpmPackage(name string) Package {
	response := Package{}

	body := Query("https://pkg.thingsima.de/npm/" + name)

	unmarshallErr := json.Unmarshal(body, &response)
	if unmarshallErr != nil {
		panic(unmarshallErr)
	}
	return response
}

// ParseRequirements takes the file path for a requirements
// file, parses it, queries each package and returns
// a slice of Package structs
func ParseNpmPackage(path string) []string {
	pkgJSON := parsers.FeastPackageJSON(path)
	sort.Strings(pkgJSON)

	return pkgJSON
}
