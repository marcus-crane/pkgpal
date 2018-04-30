package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchPypiModule(t *testing.T) {
	assert := assert.New(t)

	result := Pypi("datadog")
	assert.Equal(result.description, "The Datadog Python library")
}
