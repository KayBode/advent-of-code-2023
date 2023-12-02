package adventparser

import (
	"bufio"
	"fmt"
	"os"
)

func GetLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return make([]string, 0), err
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error on closing file %s", err)
		}
	}(f)

	result := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, nil
}
