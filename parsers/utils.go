package parsers

import (
	"io/ioutil"
	"os"
)

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

func GetCurrentDirectory() string {
	expected, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return expected
}

func FilterDirStrings(dirStrings []string) []string {
	packageStrings := []string{}
	for _, s := range dirStrings {
		switch s {
		case "package.json":
			fallthrough
		case "requirements.txt":
			packageStrings = append(packageStrings, s)
		default:
			continue
		}
	}
	return packageStrings
}

func GetDirContents(dirpath string) []string {
	fileList := []string{}
	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		fileList = append(fileList, f.Name())
	}
	return fileList
}

func DetectPackageFile(dirpath string) string {
	files := GetDirContents(dirpath)
	filteredFiles := FilterDirStrings(files)
	if len(filteredFiles) >= 1 {
		return filteredFiles[0]
	}
	return ""
}
