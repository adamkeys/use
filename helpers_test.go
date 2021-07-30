package main

import (
	"testing"
)

func TestPrefix(t *testing.T) {
	// This test assumes node is installed via homebrew.
	// TODO: Do not depend on the system environment.
	path, err := packagePrefix("node")
	if err != nil {
		t.Fatal(err)
	}
	const exp = "/usr/local/opt/node"
	if path != exp {
		t.Errorf("expected path to be: %q; got: %q", exp, path)
	}
}

func TestPackagePath(t *testing.T) {
	path := packagePath("/usr/local/opt/node", "10")
	const exp = "/usr/local/opt/node@10/bin"
	if path != exp {
		t.Errorf("expected path to be: %q; got: %q", exp, path)
	}
}

func TestJoinPath(t *testing.T) {
	cases := []struct {
		name string
		a    string
		b    string
		exp  string
	}{
		{"Left", "foo", "", "foo"},
		{"Right", "", "bar", "bar"},
		{"Both", "foo", "bar", "foo:bar"},
		{"Recursive", "foo", "bar:baz", "foo:bar:baz"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if v := joinPath(tc.a, tc.b); v != tc.exp {
				t.Errorf("expected %v; got: %v", tc.exp, v)
			}
		})
	}
}
