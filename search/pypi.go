package search

import (
	"encoding/json"
	"sort"

	"github.com/marcus-crane/pkgpal/parsers"
)

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
