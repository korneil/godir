package main

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
)

func parseArgs(args []string) (root string, patterns []*regexp.Regexp) {
	var i, j, l int

	patterns = make([]*regexp.Regexp, 0)
	for _, p := range args {
		if path.IsAbs(p) {
			root = p
			continue
		}

		l = len(p)
		for i = 0; i < l; i++ {
			for j = i; j < l && p[j] != os.PathSeparator; j++ {
			}
			if j-i == 0 {
				continue
			}
			re, err := regexp.Compile(p[i:j])
			if err != nil {
				println(err)
				os.Exit(-1)
			}
			patterns = append(patterns, re)
			i = j
		}
	}

	if root == "" {
		var err error
		root, err = os.Getwd()
		if err != nil {
			println(err)
			os.Exit(-1)
		}
	}

	return
}

func walk(root string, patterns []*regexp.Regexp) []string {
	dirs := make([]string, 1)
	dirs[0] = root
	t := make([]string, 0, 1)

	for _, re := range patterns {
		t = t[:0]
		t = append(t, dirs...)
		dirs = dirs[:0]

		for _, s := range t {
			if f, err := os.Open(s); err == nil {
				if names, err := f.Readdirnames(-1); err == nil {
					for i := range names {
						if re.Match([]byte(names[i])) {
							dirs = append(dirs, path.Join(s, names[i]))
						}
					}
				}
				_ = f.Close()
			}
		}
		if len(dirs) == 0 {
			break
		}
	}

	t = t[:0]
	for i := range dirs {
		if fi, err := os.Lstat(dirs[i]); err == nil && fi.IsDir() {
			t = append(t, dirs[i])
		}
	}

	return t
}

func main() {
	root, patterns := parseArgs(os.Args[1:])
	dirs := walk(root, patterns)

	var r string
	switch len(dirs) {
	case 0:
		os.Exit(-1)
	case 1:
		r = dirs[0]
	default:
		sort.Strings(dirs)
		for i, dir := range dirs {
			println(strconv.Itoa(i) + ": " + dir)
		}
		i := 0
		_, _ = fmt.Scanf("%d", &i)
		r = dirs[i]
	}
	println(r)
	fmt.Println(r)
}
