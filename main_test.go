package main

import (
	"regexp"
	"testing"
)

var (
	root     string
	patterns []*regexp.Regexp
	args     = []string{"/opt/local", "e"}
)

func init() {
	root, patterns = parseArgs(args)
}

func BenchmarkParseArgs(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parseArgs(args)
	}
}

func BenchmarkWalk(b *testing.B) {
	for n := 0; n < b.N; n++ {
		walk(root, patterns)
	}
}
