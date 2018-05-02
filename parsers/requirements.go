package parsers

import (
	"strings"
)

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
func PruneRequirements(requirements []string) []string {
	var reqs []string
	for _, requirement := range requirements {
		unpinnedRequirement := PruneVersion(requirement)
		reqs = append(reqs, unpinnedRequirement)
	}
	return reqs
}

// LoadRequirements loads the provided path and then parses the file it finds
func LoadRequirements(path string) []string {
	_, file := LoadFile(path, true)
	return ParseRequirements(file)
}

// FeastRequirements loads in a requirements file, parses it, breaks up each line,
// strips out versions if they're pinned and returns a list of package names
// in alphabetical order
func FeastRequirements(path string) []string {
	requirements := LoadRequirements(path)
	return PruneRequirements(requirements)
}
