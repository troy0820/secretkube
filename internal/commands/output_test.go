package commands

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutputSecret(t *testing.T) {
	t.Parallel()
	m, err := MakeMapFromJSON("../../testdata/json.json")
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
	t.Run("Secret Stringdata equals map", func(T *testing.T) {
		bytemap := TurnMapToBytes(m)
		secret, err := createSecret("fancy-secret", m, bytemap)
		if err != nil {
			t.Fatal("Error creating the secret", err)
		}
		assert.Equal(t, m, secret.StringData, "Results are not equal")
	})
}

func TestWriteToStdOut(t *testing.T) {
	t.Parallel()
	b := &bytes.Buffer{}
	m, err := MakeMapFromJSON("../../testdata/json.json")
	if err != nil {
		t.Error("Error with makeMapfromJson")
	}
	bytemap := TurnMapToBytes(m)
	secret, err := createSecret("fancy-secret", m, bytemap)
	if err != nil {
		t.Fatal(err)
	}
	if err := writeToStdOut(b, secret); err != nil {
		t.Fatal("Failure to write to stdout: ", err)
	}
}
