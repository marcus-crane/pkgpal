package parsers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFile(t *testing.T) {
	assert := assert.New(t)

	file := LoadFile("fixtures/test.txt")
	assert.Equal(file, "This is a test")
}

func TestParseRequirements(t *testing.T) {
	assert := assert.New(t)

	reqs := "amqp==2.2.2\nbeautifulsoup4==4.6.0\nbilliard==3.5.0.3"
	list := ParseRequirements(reqs)

	expected := []string{"amqp==2.2.2", "beautifulsoup4==4.6.0", "billiard==3.5.0.3"}
	assert.Equal(list, expected)
}

func TestLoadRequirements(t *testing.T) {
	assert := assert.New(t)

	file := LoadRequirements("fixtures/requirements.txt")
	assert.Equal(20, len(file))
}
