package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/fatih/color"
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

func Query(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body
}

// Pypi queries pkgparse
func Pypi(name string) Package {
	response := Package{}

	body := Query("https://pkg.thingsima.de/pypi/" + name)

	unmarshallErr := json.Unmarshal(body, &response)
	if unmarshallErr != nil {
		panic(unmarshallErr)
	}
	return response
}

func main() {
	items := parsers.LoadRequirements("requirements.txt")
	for _, item := range items {
		res := Pypi(item)
		color.Green(res.Name + " => " + res.Description)
	}
}
