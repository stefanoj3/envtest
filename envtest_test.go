package envtest_test

import (
	"encoding/hex"
	"math/rand"
	"os"
	"testing"

	"github.com/stefanoj3/envtest"
)

func TestSetEnv(t *testing.T) {
	key := randomString(t)
	val := "VAL1"

	foundValue := os.Getenv(key)
	if foundValue != "" {
		t.Errorf("expected no value for %s, got %s instead", key, foundValue)
	}
	restoreFunc := envtest.SetEnv(t, map[string]string{key: val})

	foundValue = os.Getenv(key)
	if foundValue != val {
		t.Errorf("expected %s, got %s instead", key, foundValue)
	}

	restoreFunc()
	foundValue = os.Getenv(key)
	if foundValue != "" {
		t.Errorf("expected no value for %s after the restore, got %s instead", key, foundValue)
	}
}

func randomString(t *testing.T) string {
	u := make([]byte, 16)
	_, err := rand.Read(u)
	if err != nil {
		t.Fatal("failed to generate random string: " + err.Error())
	}

	u[8] = (u[8] | 0x80) & 0xBF
	u[6] = (u[6] | 0x40) & 0x4F

	return hex.EncodeToString(u)
}
