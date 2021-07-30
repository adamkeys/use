package use_test

import (
	"testing"

	"github.com/adamkeys/use"
)

func TestPrefix(t *testing.T) {
	// This test assumes node is installed via homebrew.
	// TODO: Do not depend on the system environment.
	path, err := use.Prefix("node")
	if err != nil {
		t.Fatal(err)
	}
	const exp = "/usr/local/opt/node"
	if path != exp {
		t.Errorf("expected path to be: %q; got: %q", exp, path)
	}
}

func TestVersion(t *testing.T) {
	path, err := use.Version("/usr/local/opt/node", "10")
	if err != nil {
		t.Fatal(err)
	}
	const exp = "/usr/local/opt/node@10/bin"
	if path != exp {
		t.Errorf("expected path to be: %q; got: %q", exp, path)
	}
}
