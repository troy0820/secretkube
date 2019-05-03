package cmd

import (
	"testing"
)

func TestConverToBase64(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"hello":   {"hello", "aGVsbG8="},
		"testing": {"testing", "dGVzdGluZw=="},
	}
	for name, test := range tests {
		got := convertToBase64(test.input)
		t.Run(name, func(t *testing.T) {
			if test.want != got {
				t.Fatalf("exptected %#v, got %#v", test.want, got)
			}
		})
	}
}
