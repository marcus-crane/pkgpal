package parsers

import (
	"io/ioutil"
	"strings"
)

// LoadRequirements takes a file and loads it returning string data
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

// LoadRequirements loads the provided path and then parses the file it finds
func LoadRequirements(path string) []string {
	file := LoadFile(path)
	return ParseRequirements(file)
}
