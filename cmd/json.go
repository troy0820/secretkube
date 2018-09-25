package cmd

import (
	"bufio"
	"os"
	"strings"
)

func makeMapfromJson(file string) (map[string]interface{}, error) {
	m := map[string]interface{}{}
	f, err := os.Open(file)
	if err != nil {
		panic(err)
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
