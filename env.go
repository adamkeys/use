package use

import (
	"fmt"
	"strings"
)

// EnvironmentSet holds the environment to allow reading and manipulation.
type EnvironmentSet struct {
	environ []string
	keys    map[string]keyInfo
}

// Env returns a new EnvironmentSet for the supplied environment variables. The input is intended to be used
// by the result of the syscall.Environ() function.
func Env(environ []string) *EnvironmentSet {
	keys := make(map[string]keyInfo, len(environ))
	for i, env := range environ {
		kv := strings.Split(env, "=")
		keys[kv[0]] = keyInfo{
			index:  i,
			offset: len(kv[0]) + 1,
		}
	}
	return &EnvironmentSet{
		environ: environ,
		keys:    keys,
	}
}

// Get returns the environment variable assigned to the supplied key. An empty string is returned if the key is
// not found in the set.
func (e *EnvironmentSet) Get(key string) string {
	info, ok := e.keys[key]
	if !ok {
		return ""
	}
	return e.environ[info.index][info.offset:]
}

// Set assigns the supplied value to the associated key. If the key exists the existing value will be replaced. If
// the key does not exist, a new variable will be appended.
func (e *EnvironmentSet) Set(key, value string) {
	env := fmt.Sprintf("%s=%s", key, value)

	info, ok := e.keys[key]
	if !ok {
		e.environ = append(e.environ, env)
		e.keys[key] = keyInfo{
			index:  len(e.environ),
			offset: len(key) + 1,
		}
		return
	}
	e.environ[info.index] = env
}

// Environ returns the current environment.
func (e *EnvironmentSet) Environ() []string {
	return e.environ
}

// keyInfo holds references to the variable locations.
type keyInfo struct {
	index  int
	offset int
}
