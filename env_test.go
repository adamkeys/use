package use_test

import (
	"testing"

	"github.com/adamkeys/use"
	"github.com/google/go-cmp/cmp"
)

func TestEnv(t *testing.T) {
	env := []string{
		"PATH=/usr/local/bin:/usr/bin",
	}
	set := use.Env(env)

	t.Run("Get", func(t *testing.T) {
		if path := set.Get("PATH"); path != "/usr/local/bin:/usr/bin" {
			t.Errorf("expected PATH to be /usr/local/bin:/usr/bin; got: %s", path)
		}
		if path := set.Get("UNKNOWN"); path != "" {
			t.Errorf("expected PATH to be empty; got: %s", path)
		}
	})

	t.Run("Set", func(t *testing.T) {
		set.Set("PATH", "/usr/bin")
		if path := set.Get("PATH"); path != "/usr/bin" {
			t.Errorf("expected PATH to be /usr/bin; got: %s", path)
		}

		set.Set("PATH", "/usr/sbin")
		if path := set.Get("PATH"); path != "/usr/sbin" {
			t.Errorf("expected PATH to be /usr/sbin; got: %s", path)
		}
	})

	t.Run("Environ", func(t *testing.T) {
		if diff := cmp.Diff(set.Environ(), env); diff != "" {
			t.Error(diff)
		}
	})
}
