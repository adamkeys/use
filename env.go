package main

import (
	"strings"
)

// EnvironmentSet holds the environment to allow reading and manipulation.
type EnvironmentSet struct {
	set map[string]string
}

// newEnv returns a new EnvironmentSet for the supplied environment variables. The input is intended to be used
// by the result of the syscall.Environ() function.
func newEnv(environ []string) *EnvironmentSet {
	set := make(map[string]string, len(environ))
	for _, env := range environ {
		kv := strings.Split(env, "=")
		set[kv[0]] = kv[1]
	}
	return &EnvironmentSet{
		set: set,
	}
}

// Get returns the environment variable assigned to the supplied key. An empty string is returned if the key is
// not found in the set.
func (e *EnvironmentSet) Get(key string) string {
	return e.set[key]
}

// Set assigns the supplied value to the associated key. If the key exists the existing value will be replaced. If
// the key does not exist, a new variable will be appended.
func (e *EnvironmentSet) Set(key, value string) {
	e.set[key] = value
}

// Environ returns the current environment. Empty variables are removed.
func (e *EnvironmentSet) Environ() []string {
	env := make([]string, 0, len(e.set))
	for k, v := range e.set {
		if v == "" {
			continue
		}
		env = append(env, k+"="+v)
	}
	return env
}
