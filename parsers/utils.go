package parsers

import "io/ioutil"

// LoadFile takes a file and loads it returning string data
func LoadFile(path string, stringify bool) ([]byte, string) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	if stringify {
		return nil, string(file)
	}
	return file, ""
}
