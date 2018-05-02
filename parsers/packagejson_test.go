package parsers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePackageJSON(t *testing.T) {
	assert := assert.New(t)

	expected := PackageJSON{
		Name: "test-package",
		Dependencies: Dependency{
			"heck": "^1.2.3",
		},
		DevDependencies: Dependency{
			"nice": "0.1.2",
			"wow":  "~0.5.4",
		},
	}
	actual := ParsePackageJSON("fixtures/package.json")

	assert.Equal(expected, actual)
	assert.Equal(expected.Name, "test-package")
	assert.Equal(expected.Dependencies["heck"], "^1.2.3")
	assert.Equal(expected.DevDependencies["wow"], "~0.5.4")
}

func TestSortDependencies(t *testing.T) {
	assert := assert.New(t)

	expected := []string{"heck", "nice", "wow"}

	pkgJSON := ParsePackageJSON("fixtures/package.json")
	actual := SortDependencies(pkgJSON)

	assert.Equal(expected, actual)
}
