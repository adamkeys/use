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

	// Find the location the requested package.
	prefix, err := use.Prefix(parts[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to find package: %v", err)
		usage()
	}
	path, err := use.Version(prefix, parts[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to find version: %v", err)
		usage()
	}

	// Setup the environment and launch a shell parpared to prioritize the package.
	env := use.Env(syscall.Environ())
	env.Set("PATH", joinPath(path, env.Get("PATH")))
	env.Set("USES", joinPath(name, env.Get("USES")))
	if err = execShell(env); err != nil {
		fmt.Fprintf(os.Stderr, "failed to use %s: %v\n", name, err)
		usage()
	}
}

// usage prints application usage help information.
func usage() {
	fmt.Fprintf(os.Stderr, "usage: use package@version\n")
	os.Exit(1)
}

// execShell executes a new shell using the supplied path added to the PATH environment.
func execShell(env *use.EnvironmentSet) error {
	shell := env.Get("SHELL")
	if shell == "" {
		shell = "/bin/sh"
	}

	return syscall.Exec(shell, []string{shell}, env.Environ())
}

// joinPath concatenates a and b separated by a :.
func joinPath(a, b string) string {
	if a == "" {
		return b
	}
	if b == "" {
		return a
	}
	return a + ":" + b
}
