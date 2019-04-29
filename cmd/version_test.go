package cmd

import (
	"testing"
)

func TestVersionCommand(t *testing.T) {
	vers := "SecretKube -- version 1.0.0"
	if want, got := vers, Version("1.0.0"); want != got {
		t.Errorf("You wanted %v but got %v", want, got)
	}
}
