package ignore

import (
	"bufio"
	"github.com/gobwas/glob"
	"os"
)

type Handler struct {
	globs []glob.Glob
	parent *Handler
}

func (h *Handler) Child () Handler {
	return Handler{
		globs:  []glob.Glob{},
		parent: h,
	}
}

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

func (h *Handler) AddSingleIgnore(val string) {
	h.globs = append(h.globs, glob.MustCompile(val))
}

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
