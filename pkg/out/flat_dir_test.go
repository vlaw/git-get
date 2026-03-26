package out

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockPrintable struct {
	pathFunc     func() string
	currentFunc  func() string
	branchesFunc func() []string
}

func (m *mockPrintable) Path() string               { return m.pathFunc() }
func (m *mockPrintable) Current() string            { return m.currentFunc() }
func (m *mockPrintable) Branches() []string         { return m.branchesFunc() }
func (m *mockPrintable) BranchStatus(string) string { return "" }
func (m *mockPrintable) WorkTreeStatus() string     { return "" }
func (m *mockPrintable) Remote() string             { return "" }
func (m *mockPrintable) Errors() []string           { return nil }

func TestFlatDirPrinter_Print(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		repos    []Printable
		expected string
	}{
		{
			name:     "single repo",
			repos:    []Printable{&mockPrintable{pathFunc: func() string { return "/home/user/repos/github.com/grdl/git-get" }}},
			expected: "/home/user/repos/github.com/grdl/git-get\n",
		},
		{
			name: "multiple repos",
			repos: []Printable{
				&mockPrintable{pathFunc: func() string { return "/home/user/repos/github.com/grdl/repo1" }},
				&mockPrintable{pathFunc: func() string { return "/home/user/repos/github.com/grdl/repo2" }},
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
				&mockPrintable{pathFunc: func() string { return "/home/user/repos/github.com/grdl/repo" + string(os.PathSeparator) }},
			},
			expected: "/home/user/repos/github.com/grdl/repo\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			printer := NewFlatDirPrinter()
			result := printer.Print(tt.repos)
			assert.Equal(t, tt.expected, result)
		})
	}
}
