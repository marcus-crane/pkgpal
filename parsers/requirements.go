package parsers

import (
	"io/ioutil"
	"strings"
)

// LoadFile takes a file and loads it returning string data
func LoadFile(path string) string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(file)
}

// ParseRequirements takes a requirements.txt string and parses it into a string array
func ParseRequirements(requirements string) []string {
	return strings.Split(requirements, "\n")
}

// PruneVersion parses a requirements entry and removes any version pin if it exists
func PruneVersion(requirement string) string {
	if strings.Contains(requirement, "==") {
		return strings.Split(requirement, "==")[0]
	}
	return requirement
}

// PruneRequirements parses a slice of requirements names and removes the versioning if it exists
func PruneRequirements(requirementsString string) []string {
	requirements := ParseRequirements(requirementsString)
	var reqs []string
	for _, requirement := range requirements {
		reqs = append(reqs, PruneVersion(requirement))
	}
	return reqs
}

// LoadRequirements loads the provided path and then parses the file it finds
func LoadRequirements(path string) []string {
	file := LoadFile(path)
	return ParseRequirements(file)
}