package main

import (
	"testing"
)

var (
	root     string
	patterns []string
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
