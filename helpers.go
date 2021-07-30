package main

import (
	"bytes"
	"fmt"
	"os"
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

// packageVersion returns the path of the package for the supplied version. If the package for the specified version cannot
// be found an error will be returned.
func packageVersion(prefix, version string) (string, error) {
	builder := strings.Builder{}
	builder.WriteString(prefix)
	builder.WriteByte('@')
	builder.WriteString(version)
	builder.WriteString("/bin")
	path := builder.String()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", fmt.Errorf("%s: %w", path, err)
	}
	return path, nil
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
