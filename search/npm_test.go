package search

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func TestNpmPackage(t *testing.T) {
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

	httpmock.RegisterResponder("GET", "https://pkg.thingsima.de/npm/fakehotdogpackage",
		httpmock.NewStringResponder(200, string(msg)))

	actual := NpmPackage("fakehotdogpackage")

	assert.Equal(expected, actual)
}

func TestParseNpmPackage(t *testing.T) {
	assert := assert.New(t)

	var pkgResponse Packages

	buttcli := Package{
		Description:   "The premiere farting tool",
		Homepage:      "https://butt.cli",
		LatestVersion: "0.9.9",
		Name:          "buttcli",
		PackagePage:   "https://pypi.org/project/buttcli",
		SourceRepo:    false,
		Tarball:       "https://butts.net/fart.zip",
	}

	fake5000 := Package{
		Description:   "A chatbot that is made of cardboard",
		Homepage:      "https://notreal.net",
		LatestVersion: "0.1.2",
		Name:          "fake5000",
		PackagePage:   "https://pypi.org/project/fake5000",
		SourceRepo:    false,
		Tarball:       "https://cardboarddreams.net/sigh.7z",
	}

	testpackage := Package{
		Description:   "Gotta test 'em all",
		Homepage:      "https://test.se",
		LatestVersion: "0.1.2",
		Name:          "testpackage",
		PackagePage:   "https://pypi.org/project/testpackage",
		SourceRepo:    false,
		Tarball:       "https://air.drop/isitonline.rar",
	}

	pkgResponse = append(pkgResponse, buttcli)
	pkgResponse = append(pkgResponse, fake5000)
	pkgResponse = append(pkgResponse, testpackage)

	butt, err := json.Marshal(buttcli)
	if err != nil {
		panic(err)
	}

	fake, err := json.Marshal(fake5000)
	if err != nil {
		panic(err)
	}

	test, err := json.Marshal(testpackage)
	if err != nil {
		panic(err)
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://pkg.thingsima.de/npm/buttcli",
		httpmock.NewStringResponder(200, string(butt)))

	httpmock.RegisterResponder("GET", "https://pkg.thingsima.de/npm/fake5000",
		httpmock.NewStringResponder(200, string(fake)))

	httpmock.RegisterResponder("GET", "https://pkg.thingsima.de/npm/testpackage",
		httpmock.NewStringResponder(200, string(test)))

	pkgJSON := ParseNpmPackage("fixtures/package.json")

	assert.Equal(pkgJSON, pkgResponse)
}
