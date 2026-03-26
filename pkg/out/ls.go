package out

import (
	"os"
	"strings"
)

type LsPrinter struct{}

func NewLsPrinter() *LsPrinter {
	return &LsPrinter{}
}

func (p *LsPrinter) Print(repos []Printable) string {
	var str strings.Builder

	for _, repo := range repos {
		str.WriteString(strings.TrimSuffix(repo.Path(), string(os.PathSeparator)))
		str.WriteString("\n")
	}

	return str.String()
}
