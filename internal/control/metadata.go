package control

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

const metadataFile = "metadata.yaml"

// Metadata represent the package metadata
type Metadata struct {
	// Package is the control package name
	Package string
	// The Go import path
	ImportPath string
	// List of the package maintainers
	// i.e who take the responsibility for uploading & managing it
	Maintainers []string
	// List of the packages built by this control package
	Packages []Package
}

// Package represent a package installable
type Package struct {
	// The package name
	Package string
	// Main is the full path to the file containing the `func main()` sequence
	Main string `yaml:"main,omitempty"`
	// Human description of the package
	Description string
	// Targets describe the build target (os,arches)
	Targets map[string][]string `yaml:"targets,omitempty"`
}

// IsSource returns true if the package is a source package
func (p Package) IsSource() bool {
	return p.Main == ""
}

// writeMetadata write the given metadata
func writeMetadata(m Metadata, path string) error {
	b, err := yaml.Marshal(m)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fmt.Sprintf("%s/%s", path, metadataFile), b, 0640)
}

// ReadMetadata read metadata from file
func readMetadata(path string) (Metadata, error) {
	var m Metadata

	f, err := os.Open(filepath.Join(path, metadataFile))
	if err != nil {
		return Metadata{}, err
	}
	defer f.Close()

	if err := yaml.NewDecoder(f).Decode(&m); err != nil {
		return Metadata{}, err
	}

	return m, nil
}
