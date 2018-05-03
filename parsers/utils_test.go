package parsers

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFileAsString(t *testing.T) {
	assert := assert.New(t)

	_, file := LoadFile("fixtures/test.txt", true)
	expected := ""

	assert.Equal(reflect.TypeOf(file), reflect.TypeOf(expected))
	assert.Equal(file, "This is a test")
}

func TestLoadFileAsBytes(t *testing.T) {
	assert := assert.New(t)

	file, _ := LoadFile("fixtures/test.txt", false)
	expected := []byte{}

	assert.Equal(reflect.TypeOf(file), reflect.TypeOf(expected))
}

func TestGetCurrentDirectory(t *testing.T) {
	assert := assert.New(t)

	expected, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	actual := GetCurrentDirectory()

	assert.Equal(expected, actual)
}

func TestFilterDirStrings(t *testing.T) {
	assert := assert.New(t)

	expected := []string{
		"package.json",
		"requirements.txt",
	}

	dirStrings := []string{
		"package.json",
		"requirements.txt",
		"test.txt",
	}

	actual := FilterDirStrings(dirStrings)

	assert.Equal(expected, actual)
}

func TestGetDirContents(t *testing.T) {
	assert := assert.New(t)

	expected := []string{
		".DS_Store",
		"fixtures",
		"packagejson.go",
		"packagejson_test.go",
		"requirements.go",
		"requirements_test.go",
		"utils.go",
		"utils_test.go",
	}

	cwd := GetCurrentDirectory()
	actual := GetDirContents(cwd)

	assert.Equal(expected, actual)
}

func TestDetectPackageFilesProvidedDir(t *testing.T) {
	assert := assert.New(t)

	expected := "package.json"

	cwd := GetCurrentDirectory() + "/fixtures"
	actual := DetectPackageFile(cwd)

	assert.Equal(expected, actual)
}
