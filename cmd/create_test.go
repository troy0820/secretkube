package cmd

import (
	"testing"
)

func TestConverToBase64(t *testing.T) {
	hello := convertToBase64("hello")

	if hello != "aGVsbG8=" {
		t.Errorf("Exptected should be aGVsbG8= but got %s", hello)
	}
}
