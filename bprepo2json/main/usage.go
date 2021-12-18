package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	// main operation mode
	repoList   = flag.String("l", "framework.repo.list", "list repos' name")
	outputFile = flag.String("o", "repo_pkg_module.json", "output/path/file.ext")
	exitCode   = 0
)

func usage() {
	usageViolation("")
}

func usageViolation(violation string) {
	fmt.Fprintln(os.Stderr, violation)
	fmt.Fprintln(os.Stderr, "usage: bpfmt [flags] [path ...]")
	flag.PrintDefaults()
	os.Exit(2)
}

func report(err error) {
	fmt.Fprintln(os.Stderr, err)
	exitCode = 2
}
