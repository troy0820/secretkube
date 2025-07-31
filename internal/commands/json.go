package commands

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func makeMapfromJson(file string) (map[string]any, error) {
	m := map[string]any{}
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

// MakeMapFromJSON opens a JSON file and creates a map[string]string
func MakeMapFromJSON(file string) (map[string]string, error) {
	m := map[string]any{}
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var s string
	for scanner.Scan() {
		s += scanner.Text()
	}

	if err := json.Unmarshal([]byte(s), &m); err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}
	mm := map[string]string{}
	for k, v := range m {
		switch x := v.(type) {
		case string:
			mm[k] = x
		case float64:
			if x == float64(int64(x)) {
				mm[k] = fmt.Sprintf("%.0f", x)
			} else {
				mm[k] = fmt.Sprintf("%f", x)
			}
		case bool:
			mm[k] = fmt.Sprintf("%v", x)

		}
	}
	return mm, nil
}

func stripQuotesforSecret(m map[string]string) map[string]string {
	newMap := map[string]string{}
	for k, v := range m {
		if unicode.IsDigit(rune(v[0])) || unicode.IsLetter(rune(v[0])) {
			newMap[k[1:len(k)-1]] = v[0 : len(v)-1]
		} else {
			newMap[k[1:len(k)-1]] = v[1 : len(v)-2]
		}
	}
	return newMap
}

func turnMaptoBytes(m map[string]any) map[string][]byte {
	newMap := map[string][]byte{}
	for k, v := range m {
		a := v.(string)
		if strings.ContainsAny(a, ",") && strings.ContainsAny(a, "\"") {
			newMap[k[1:len(k)-1]] = []byte(a[1 : len(a)-2])
		} else if strings.ContainsAny(a, "\"") {
			newMap[k[1:len(k)-1]] = []byte(a[1 : len(a)-1])
		} else if strings.ContainsAny(a, ",") {
			newMap[k[1:len(k)-1]] = []byte(a[0 : len(a)-1])
		}

	}
	return newMap
}

// TurnMapToBytes takes map[string]string and transforms the value to a slice of bytes
func TurnMapToBytes(m map[string]string) map[string][]byte {
	newMap := map[string][]byte{}
	for k, v := range m {
		newMap[k] = []byte(v)
	}
	return newMap
}

func convertMapValuesToBase64(m map[string][]byte) {
	for k, v := range m {
		m[k] = []byte(base64.StdEncoding.EncodeToString(v))
	}

}

func turnMaptoString(m map[string]any) map[string]string {
	newMap := map[string]string{}
	for k, v := range m {
		newMap[k] = v.(string)
	}
	return newMap
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
