package cmd

import (
	"reflect"
	"testing"
)

func TestOutputSecret(t *testing.T) {
	m, err := MakeMapFromJSON("../json.json")
	if err != nil {
		t.Error("Error with makeMapfromJson")
	}
	t.Run("Created Secrets are equal", func(t *testing.T) {
		bytemap := TurnMapToBytes(m)
		secret, err := CreateSecret("fancy-secret", bytemap)
		if err != nil {
			t.Error("Error with creating secret")
		}
		if want, got := "fancy-secret", secret.Name; want != got {
			t.Errorf("Secret Name is %v and you got %v", want, got)
		}
		if want, got := "v1", secret.APIVersion; want != got {
			t.Errorf("Secret APIVersion is %v and you got %v", want, got)
		}
		if want, got := bytemap, secret.Data; reflect.DeepEqual(want, got) != true {
			t.Errorf("Secret Data is %v and you got %v", want, got)
		}
		if want, got := "Secret", secret.Kind; want != got {
			t.Errorf("Secret Kind is %v and you got %v", want, got)
		}
	})
	t.Run("Output secret equals stdOut", func(T *testing.T) {
		//TODO: Compare output of createoutputsecret to output template
	})
}
