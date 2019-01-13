package envtest

import (
	"os"
	"testing"
)

// SetEnv set the env variables using the key-value pairs provided
// It returns a closure that can be called to restore the variables to the previous state
func SetEnv(t *testing.T, keyValuePairs map[string]string) func() {
	toRestore := map[string]string{}

	for key, value := range keyValuePairs {
		toRestore[key] = os.Getenv(key)
		err := os.Setenv(key, value)
		if err != nil {
			t.Errorf("envtest: failed to set env variable(%s): %s", key, err.Error())
		}
	}

	return func() {
		for key, value := range toRestore {
			err := os.Setenv(key, value)
			if err != nil {
				t.Errorf("envtest: failed to restore env variable(%s): %s", key, err.Error())
			}
		}
	}
}
