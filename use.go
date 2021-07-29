// Package use provides functions to use a version of a brew package.
package use

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Prefix returns the result of brew --prefix <packageName>
func Prefix(packageName string) (string, error) {
	buf := bytes.Buffer{}
	cmd := exec.Command("brew", "--prefix", packageName)
	cmd.Stdout = &buf
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("run: %w", err)
	}

	path := bytes.TrimRight(buf.Bytes(), "\r\n")
	return string(path), nil
}

// Version returns the path of the package for the supplied version. If the package for the specified version cannot
// be found an error will be returned.
func Version(packageName, version string) (string, error) {
	prefix, err := Prefix(packageName)
	if err != nil {
		return "", err
	}

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
