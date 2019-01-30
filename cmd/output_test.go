package cmd

import "testing"

func TestOutputSecret(t *testing.T) {
	m, err := makeMapfromJson("../json.json")
	if err != nil {
		t.Error("Error with makeMapfromJson")
	}
	t.Run("Created Secrets are equal", func(t *testing.T) {

	})
	t.Run("Output secret equals stdOut", func(T *testing.T) {

	})
	print(m)
}
