package cmd

import (
	"bytes"
	"testing"
)

func TestVersionCommand(t *testing.T) {
	vers := "SecretKube -- version 1.0.0"
	buffer := bytes.NewBufferString(vers)
	if want, got := buffer.String(), Version("1.0.0"); want != got {
		t.Errorf("You wanted %v but got %v", want, got)
	}
}
