package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	fPath := flag.String("path", "", "Filepath for input")
	flag.Parse()

	lines, err := getInput(*fPath)
	if err != nil {
		return err
	}

	var result int
	for _, line := range lines {
		firstNumber, err := findFirstNumber(line)
		if err != nil {
			fmt.Println(err)
			continue
		}

		lastNumber, _ := findLastNumber(line)

		combinedNumber := fmt.Sprintf("%s%s", firstNumber, lastNumber)
		number, err := strconv.Atoi(combinedNumber)
		if err != nil {
			return err
		}

		result = result + number
	}

	_, _ = fmt.Fprintf(os.Stdout, "Sum is: %d", result)

	return nil
}

func extractNumberFromString(value string) string {
	re := regexp.MustCompile(`[0-9]+|one|two|three|four|five|six|seven|eight|nine`)
	result := re.FindString(value)
	return result
}

func transformTextToNumber(value string) string {
	switch value {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return value
	}
}

func findFirstNumber(line string) (string, error) {
	for i := 0; i <= len(line); i++ {
		subString := line[:i]
		number := extractNumberFromString(subString)
		if len(number) != 0 {
			result := transformTextToNumber(number)
			return result, nil
		}
	}

	return "", errors.New("no number found in line")
}

func findLastNumber(line string) (string, error) {
	for i := len(line); i >= 0; i-- {
		subString := line[i:]
		number := extractNumberFromString(subString)
		if len(number) != 0 {
			result := transformTextToNumber(number)
			return result, nil
		}
	}

	return "", errors.New("no number found in line")
}

func getInput(path string) ([]string, error) {
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
