package out

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockPrintableLs struct {
	pathFunc     func() string
	currentFunc  func() string
	branchesFunc func() []string
}

func (m *mockPrintableLs) Path() string               { return m.pathFunc() }
func (m *mockPrintableLs) Current() string            { return m.currentFunc() }
func (m *mockPrintableLs) Branches() []string         { return m.branchesFunc() }
func (m *mockPrintableLs) BranchStatus(string) string { return "" }
func (m *mockPrintableLs) WorkTreeStatus() string     { return "" }
func (m *mockPrintableLs) Remote() string             { return "" }
func (m *mockPrintableLs) Errors() []string           { return nil }

func TestLsPrinter_Print(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		repos    []Printable
		expected string
	}{
		{
			name:     "single repo",
			repos:    []Printable{&mockPrintableLs{pathFunc: func() string { return "/home/user/repos/github.com/grdl/git-get" }}},
			expected: "/home/user/repos/github.com/grdl/git-get\n",
		},
		{
			name: "multiple repos",
			repos: []Printable{
				&mockPrintableLs{pathFunc: func() string { return "/home/user/repos/github.com/grdl/repo1" }},
				&mockPrintableLs{pathFunc: func() string { return "/home/user/repos/github.com/grdl/repo2" }},
			},
			expected: "/home/user/repos/github.com/grdl/repo1\n/home/user/repos/github.com/grdl/repo2\n",
		},
		{
			name:     "empty list",
			repos:    []Printable{},
			expected: "",
		},
		{
			name: "path with trailing separator",
			repos: []Printable{
				&mockPrintableLs{pathFunc: func() string { return "/home/user/repos/github.com/grdl/repo" + string(os.PathSeparator) }},
			},
			expected: "/home/user/repos/github.com/grdl/repo\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			printer := NewLsPrinter()
			result := printer.Print(tt.repos)
			assert.Equal(t, tt.expected, result)
		})
	}
}
