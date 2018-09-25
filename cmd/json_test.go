package cmd

import "testing"

func TestJsonFunction(t *testing.T) {
	m, err := makeMapfromJson("../json.json")

	if err != nil {
		t.Error("Error when executing function")
	}
	for k, _ := range m {
		if key, ok := m[k]; ok {
			continue
		} else {
			t.Error("Error: Test failed retrieving key:", key)
		}
	}

}
