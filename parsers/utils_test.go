package parsers

import (
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
