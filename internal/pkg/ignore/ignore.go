package ignore

import (
	"bufio"
	"github.com/gobwas/glob"
	"os"
)

// Handler takes care of the ignore logic.
type Handler struct {
	globs []glob.Glob
	parent *Handler
}

// Child creates a child handler which will use it's own tests then pass the file to the
// parent and allow parents in a chain upwards to test the file in question.
func (h *Handler) Child () Handler {
	return Handler{
		globs:  []glob.Glob{},
		parent: h,
	}
}

// AddIgnoreSource adds a file as a ignore source.
// This file will read and add each row of the file as a glob entry to the
// handlers ignore list.
func (h *Handler) AddIgnoreSource(source string) error {
	file, err := os.Open(source)
	if err != nil {
		return err
	}

	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		h.globs = append(h.globs, glob.MustCompile(scan.Text()))
	}

	if err := scan.Err(); err != nil {
		return err
	}

	return nil
}

// AddSingleIgnore adds a single glob value to the list of ignored files.
func (h *Handler) AddSingleIgnore(val string) {
	h.globs = append(h.globs, glob.MustCompile(val))
}

// IsIgnored checks if a file is ignored with the handlers ruleset.
func (h *Handler) IsIgnored(file string) bool {
	for _, g := range h.globs {
		if g.Match(file) {
			return true
		}
	}

	if h.parent == nil  {
		return false
	}
	return  h.parent.IsIgnored(file)
}
