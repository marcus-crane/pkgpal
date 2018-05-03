package search

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/marcus-crane/pkgpal/parsers"
)

// PypiPackage queries the /pypi/<package> endpoint for Python packages
func PypiPackage(name string) Package {
	response := Package{}

	body := Query("https://pkg.thingsima.de/pypi/" + name)

	unmarshallErr := json.Unmarshal(body, &response)
	if unmarshallErr != nil {
		fmt.Println("Paniced at unmarshall")
		panic(unmarshallErr)
	}
	return response
}

// ParseRequirements takes the file path for a requirements
// file, parses it, queries each package and returns
// a slice of Package structs
func ParseRequirements(path string) []string {
	requirements := parsers.FeastRequirements(path)
	sort.Strings(requirements)

	return requirements
}
