package search

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func TestPypiPackage(t *testing.T) {
	assert := assert.New(t)

	expected := Package{
		Description:   "The hotdog library",
		Homepage:      "https://hotdog.app",
		LatestVersion: "0.1.2",
		Name:          "fakehotdogpackage",
		PackagePage:   "https://pypi.org/project/notarealthing",
		SourceRepo:    false,
		Tarball:       "https://src.things.net/hotdog/0.1.2/src.tar.gz",
	}

	msg, err := json.Marshal(expected)
	if err != nil {
		panic(err)
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://pkg.thingsima.de/pypi/fakehotdogpackage",
		httpmock.NewStringResponder(200, string(msg)))

	actual := PypiPackage("fakehotdogpackage")

	assert.Equal(expected, actual)
}
