package main

import (
	"advent-of-code-2023/internal/adventparser"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type game struct {
	id          int
	colorCounts []colorCount
}

type colorCount struct {
	color string
	count int
}

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	fPath := flag.String("path", "", "Filepath for input")
	flag.Parse()

	lines, err := adventparser.GetLines(*fPath)
	if err != nil {
		return err
	}

	var result int
	for _, line := range lines {
		g, err := createGameFromLine(line)
		if err != nil {
			return err
		}

		red, blue, green := getLowestNeededColorCount(g)

		result = result + red*blue*green
	}

	_, _ = fmt.Fprintf(os.Stdout, "Result is: %d\n", result)

	return nil
}

func createGameFromLine(line string) (game, error) {
	re := regexp.MustCompile(`[0-9]+`)
	idString := re.FindString(line)
	id, err := strconv.Atoi(idString)
	if err != nil {
		return game{}, err
	}

	idIndex := strings.Index(line, ":")
	gameString := line[idIndex+2:]

	g := game{
		id:          id,
		colorCounts: make([]colorCount, 0),
	}
	splits := make([]string, 0)
	r := strings.Split(gameString, "; ")
	for _, s := range r {
		x := strings.Split(s, ", ")
		splits = append(splits, x...)
	}

	for _, split := range splits {
		round := strings.Split(split, " ")
		c, err := strconv.Atoi(round[0])
		if err != nil {
			return game{}, err
		}

		cCount := colorCount{color: round[1], count: c}

		g.colorCounts = append(g.colorCounts, cCount)
	}

	return g, nil
}

func getLowestNeededColorCount(g game) (red int, blue int, green int) {
	for _, cCount := range g.colorCounts {
		switch cCount.color {
		case "red":
			if cCount.count > red {
				red = cCount.count
			}
		case "blue":
			if cCount.count > blue {
				blue = cCount.count
			}
		case "green":
			if cCount.count > green {
				green = cCount.count
			}
		}
	}

	return red, blue, green
}

func checkGame(g game, red int, blue int, green int) bool {
	for _, cCount := range g.colorCounts {
		var checkNumber int

		switch cCount.color {
		case "red":
			checkNumber = red
		case "blue":
			checkNumber = blue
		case "green":
			checkNumber = green
		}

		if cCount.count > checkNumber {
			return false
		}
	}

	return true
}
