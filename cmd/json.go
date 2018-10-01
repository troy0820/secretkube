package cmd

import (
	"bufio"
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
			m[str2] = string(str[1])
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
