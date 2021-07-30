package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// packagePrefix returns the result of brew --prefix <packageName>
func packagePrefix(packageName string) (string, error) {
	buf := bytes.Buffer{}
	cmd := exec.Command("brew", "--prefix", packageName)
	cmd.Stdout = &buf
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("run: %w", err)
	}

	path := bytes.TrimRight(buf.Bytes(), "\r\n")
	return string(path), nil
}

// packagePath returns the path of the package for the supplied version. If the package for the specified version cannot
// be found an error will be returned.
func packagePath(prefix, version string) string {
	builder := strings.Builder{}
	builder.WriteString(prefix)
	builder.WriteByte('@')
	builder.WriteString(version)
	builder.WriteString("/bin")
	return builder.String()
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
