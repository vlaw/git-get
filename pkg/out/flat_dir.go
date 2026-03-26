package out

import (
	"os"
	"strings"
)

// FlatDirPrinter prints a list of repos in a flat-dir format.
type FlatDirPrinter struct{}

// NewFlatDirPrinter creates a FlatDirPrinter.
func NewFlatDirPrinter() *FlatDirPrinter {
	return &FlatDirPrinter{}
}

// Print generates a flat list of repository directory absolute paths - each path on a new line.
func (p *FlatDirPrinter) Print(repos []Printable) string {
	var str strings.Builder

	for _, repo := range repos {
		str.WriteString(strings.TrimSuffix(repo.Path(), string(os.PathSeparator)))
		str.WriteString("\n")
	}

	return str.String()
}
