package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/adamkeys/use"
)

func main() {
	// Determine the requested package and version. The user specifies their choice by suppling a single command
	// line argument in the format package@version.
	if len(os.Args) < 2 {
		usage()
	}
	name := os.Args[1]
	parts := strings.Split(name, "@")
	if len(parts) != 2 {
		usage()
	}

	path, err := use.Version(parts[0], parts[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to find package: %v", err)
		usage()
	}

	if err = execShell(name, path); err != nil {
		fmt.Fprintf(os.Stderr, "failed to use %s: %v\n", name, err)
		os.Exit(1)
	}
}

// usage prints application usage help information.
func usage() {
	fmt.Fprintf(os.Stderr, "usage: use package@version\n")
	os.Exit(1)
}

// execShell executes a new shell using the supplied path added to the PATH environment.
func execShell(name, path string) error {
	env := use.Env(syscall.Environ())
	envPath := env.Get("PATH")
	env.Set("PATH", path+":"+envPath)

	shell := env.Get("SHELL")
	if shell == "" {
		shell = "/bin/sh"
	}

	return syscall.Exec(shell, []string{shell}, env.Environ())
}
