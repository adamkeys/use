package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEnv(t *testing.T) {
	env := []string{
		"PATH=/usr/local/bin:/usr/bin",
	}
	set := newEnv(env)

	t.Run("Get", func(t *testing.T) {
		cases := []struct {
			key   string
			value string
		}{
			{"PATH", "/usr/local/bin:/usr/bin"},
			{"UNKNOWN", ""},
		}
		for _, tc := range cases {
			t.Run(tc.key, func(t *testing.T) {
				if v := set.Get(tc.key); v != tc.value {
					t.Errorf("expected %s to be %v; got: %s", tc.key, tc.value, v)
				}
			})
		}
	})

	t.Run("Set", func(t *testing.T) {
		cases := []struct {
			key   string
			value string
		}{
			{"PATH", "/usr/bin"},
			{"PATH", "/usr/sbin"},
			{"USES", "node@12"},
			{"USES", "node@12"},
		}
		for _, tc := range cases {
			t.Run(tc.key, func(t *testing.T) {
				set.Set(tc.key, tc.value)
				if v := set.Get(tc.key); v != tc.value {
					t.Errorf("expected %s to be %s; got: %s", tc.key, tc.value, v)
				}
			})
		}
	})

	t.Run("Environ", func(t *testing.T) {
		set := newEnv(env)
		set.Set("USES", "")
		if diff := cmp.Diff(set.Environ(), env); diff != "" {
			t.Error(diff)
		}
	})
}
