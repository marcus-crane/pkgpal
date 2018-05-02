package search

import (
	"encoding/json"
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
