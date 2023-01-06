package initializers

import (
	"os"
	"testing"
)

func TestEnvFile(t *testing.T) {
	prepareEnvFile()

	got := os.Getenv("TEST")
	want := "ok"

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func prepareEnvFile() {
	EnvironmentVariables()
}
