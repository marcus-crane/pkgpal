package parsers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRequirements(t *testing.T) {
	assert := assert.New(t)

	reqs := "amqp==2.2.2\nbeautifulsoup4==4.6.0\nbilliard==3.5.0.3"
	list := ParseRequirements(reqs)

	expected := []string{"amqp==2.2.2", "beautifulsoup4==4.6.0", "billiard==3.5.0.3"}
	assert.Equal(list, expected)
}

func TestPruneVersion(t *testing.T) {
	assert := assert.New(t)

	req := "amqp==2.2.2"
	actual := PruneVersion(req)

	expected := "amqp"
	assert.Equal(actual, expected)
}

func TestPruneVersionMissingVersion(t *testing.T) {
	assert := assert.New(t)

	req := "amqp"
	actual := PruneVersion(req)

	expected := "amqp"
	assert.Equal(actual, expected)
}

func TestPruneVersions(t *testing.T) {
	assert := assert.New(t)

	reqs := []string{"amqp==2.2.2", "beautifulsoup4==4.6.0", "billiard==3.5.0.3"}
	actual := PruneRequirements(reqs)

	expected := []string{"amqp", "beautifulsoup4", "billiard"}
	assert.Equal(expected, actual)
}

func TestPruneMixedVersions(t *testing.T) {
	assert := assert.New(t)

	reqs := []string{"amqp", "beautifulsoup4==4.6.0", "awscli==0.1.2.3.4.5", "python"}
	actual := PruneRequirements(reqs)

	expected := []string{"amqp", "beautifulsoup4", "awscli", "python"}
	assert.Equal(expected, actual)
}

func TestLoadRequirements(t *testing.T) {
	assert := assert.New(t)

	file := LoadRequirements("fixtures/requirements.txt")
	assert.Equal(20, len(file))
}

func TestFeastRequirements(t *testing.T) {
	assert := assert.New(t)

	expected := []string{
		"amqp", "beautifulsoup4", "billiard", "celery", "certifi", "chardet",
		"django", "gunicorn", "httplib2", "idna", "kombu", "mistune-contrib",
		"mistune", "psycopg2-binary", "pygments", "pytvdbapi", "pytz",
		"requests", "urllib3", "vine",
	}
	actual := FeastRequirements("fixtures/requirements.txt")

	assert.Equal(expected, actual)
}
