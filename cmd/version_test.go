package cmd

import (
	"strings"
	"testing"
)

func TestVersionCommand(t *testing.T) {
	vers := "1.0.0"
	version := strings.ContainsAny(Version(vers), "1.0.0")
	if version == false {
		t.Errorf("Expected string should have %s", vers)
	}
}
