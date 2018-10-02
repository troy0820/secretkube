package cmd

import (
	"reflect"
	"testing"
)

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

func TestMapToString(t *testing.T) {
	m, err := makeMapfromJson("../json.json")

	if err != nil {
		t.Error("Error when executing function")
	}

	mm := turnMaptoString(m)

	for _, v := range mm {
		if reflect.TypeOf(v).Kind() == reflect.String {
			continue
		}
	}
}

func TestMapToBytes(t *testing.T) {
	m, err := makeMapfromJson("../json.json")

	if err != nil {
		t.Error("Error when executing function")
	}

	mm := turnMaptoBytes(m)

	for _, v := range mm {
		if reflect.TypeOf(v).Kind() == reflect.String {
			continue
		}
	}
}

func TestConvertMapToBase64(t *testing.T) {
	m, err := makeMapfromJson("../json.json")

	if err != nil {
		t.Error("Error when executing function")
	}

	mm := convertMapValuesToBase64(turnMaptoBytes(m))

	for _, v := range mm {
		if reflect.TypeOf(v).Kind() == reflect.Uint8 {
			continue
		}

	}
}
