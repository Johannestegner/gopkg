package control

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const changelogFile = "changelog.yaml"

// Changelog is the root object containing all package releases
type Changelog struct {
	Releases []Release
}

// Release is produced each time a package is released
type Release struct {
	// The package version number (upstream-internal)
	// f.e 1.2.0-1 is the initial release of upstream version 1.2.0.
	Version string
	// Who has taking care of the release upload
	Uploader string
	// The human descriptions of changes applied since last release
	Changes []string
}

// NewChangelog create a brand new changelog
func NewChangelog(initialVersion, uploader string) Changelog {
	return Changelog{
		Releases: []Release{{
			Version:  fmt.Sprintf("%s-1", initialVersion),
			Uploader: uploader,
			Changes:  []string{"Initial release"},
		}},
	}
}

// WriteChangelog write the given changelog
func WriteChangelog(c Changelog, path string) error {
	b, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fmt.Sprintf("%s/%s", path, changelogFile), b, 0640)
}
