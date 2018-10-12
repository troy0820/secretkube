package cmd

import (
	"reflect"
	"testing"
	"unicode"
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
		if reflect.TypeOf(v[0]).Kind() != reflect.Uint8 {
			t.Errorf("Error: Value not a byte slice but it's a  %T", v)
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
		if reflect.TypeOf(v[0]).Kind() != reflect.Uint8 {
			t.Errorf("Error: Value is not a byte slice but it's a %T", v)
		}

	}
}

func TestStripQuotes(t *testing.T) {
	m, err := makeMapfromJson("../json.json")

	if err != nil {
		t.Error("Error executing function")
	}
	newMap := stripQuotesforSecret(turnMaptoString(m))
	for _, v := range newMap {
		if unicode.IsDigit(rune(v[0])) || unicode.IsLetter(rune(v[0])) || unicode.IsSymbol(rune(v[0])) {
			continue
		} else {
			t.Error("Error: Quotes still remain in map", v)
		}
	}
}
