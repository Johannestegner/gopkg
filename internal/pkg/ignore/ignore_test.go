package ignore

import (
	"io/ioutil"
	"testing"
)

func TestHandler_AddIgnoreSource(t *testing.T) {
	dir, _ := ioutil.TempDir("", "*")
	file, _ := ioutil.TempFile(dir, "ignore")

	file.WriteString("test.txt\n")
	file.WriteString("*.json\n")
	file.WriteString(".env")
	file.Close()

	ignore := Handler{}

	ignore.AddIgnoreSource(file.Name())

	if !ignore.IsIgnored("test.txt") {
		t.Errorf("direct hit was not ignored")
	}

	if !ignore.IsIgnored("something.json") {
		t.Errorf("globbed single was not ignored")
	}

	if !ignore.IsIgnored(".env") {
		t.Errorf("env not ignored")
	}

	if ignore.IsIgnored(".envs") {
		t.Errorf("envs was ignored")
	}
}

func TestHandler_AddSingleIgnore(t *testing.T) {
	ignore := Handler{}

	ignore.AddSingleIgnore(".env")

	if !ignore.IsIgnored(".env") {
		t.Errorf("env not ignored")
	}

	if ignore.IsIgnored(".envs") {
		t.Errorf("envs was ignored")
	}
}

func TestHandler_Child(t *testing.T) {
	ignore := Handler{}
	ignore.AddSingleIgnore(".env")

	child := ignore.Child()
	child.AddSingleIgnore("test.txt")

	if ignore.IsIgnored("test.txt") {
		t.Error("parent uses child rules")
	}

	if !ignore.IsIgnored(".env") {
		t.Error("env was not ignored")
	}

	if !child.IsIgnored("test.txt") {
		t.Error("test.txt was not ignored")
	}

	if !child.IsIgnored(".env") {
		t.Error("child does not use parent rules")
	}

}
