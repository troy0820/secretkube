package cmd

import (
	"bufio"
	"encoding/base64"
	"errors"
	"os"
	"strings"
)

func makeMapfromJson(file string) (map[string]interface{}, error) {
	m := map[string]interface{}{}
	f, err := os.Open(file)
	if err != nil {
		return m, errors.New(err.Error())
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		str := strings.SplitN(scanner.Text(), ":", 2)
		if len(str) > 1 {
			str2 := strings.Trim(str[0], " ")
			m[str2] = string(strings.Trim(str[1], " "))
		}
	}
	return m, nil
}

func turnMaptoBytes(m map[string]interface{}) map[string][]byte {
	newMap := map[string][]byte{}
	for k, v := range m {
		newMap[k] = []byte(v.(string))
	}
	return newMap
}

func convertMapValuesToBase64(m map[string][]byte) map[string][]byte {
	newMap := map[string][]byte{}
	for k, v := range m {
		newMap[k] = []byte(base64.StdEncoding.EncodeToString(v))
	}
	return newMap
}

func turnMaptoString(m map[string]interface{}) map[string]string {
	newMap := map[string]string{}
	for k, v := range m {
		newMap[k] = v.(string)
	}
	return newMap
}

func decodeFromBase64(str string) string {
	str2, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		os.Exit(1)
	}
	return string(str2)
}

func saveToFile(str, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString(str)
	w.Flush()
}
