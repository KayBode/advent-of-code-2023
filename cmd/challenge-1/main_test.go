package main

import (
	"testing"
)

func TestFirstNumber(t *testing.T) {
	line1 := "five4sevenvlhnczltpxtskstcxqmfkstgnine"
	line2 := "threeight"
	line3 := "9"
	line4 := "one"

	expected1 := "5"
	expected2 := "3"
	expected3 := "9"
	expected4 := "1"

	actual1, _ := findFirstNumber(line1)
	actual2, _ := findFirstNumber(line2)
	actual3, _ := findFirstNumber(line3)
	actual4, _ := findFirstNumber(line4)

	if actual1 != expected1 {
		t.Fatalf("Acutal 1 was: %s; Expected was %s", actual1, expected1)
	}

	if actual2 != expected2 {
		t.Fatalf("Acutal 2 was: %s; Expected was %s", actual2, expected2)
	}

	if actual3 != expected3 {
		t.Fatalf("Acutal 3 was: %s; Expected was %s", actual3, expected3)
	}

	if actual4 != expected4 {
		t.Fatalf("Acutal 4 was: %s; Expected was %s", actual4, expected4)
	}
}

func TestLastNumber(t *testing.T) {
	line1 := "five4sevenvlhnczltpxtskstcxqmfkstgnine"
	line2 := "threeight"
	line3 := "9"
	line4 := "one"

	expected1 := "9"
	expected2 := "8"
	expected3 := "9"
	expected4 := "1"

	actual1, _ := findLastNumber(line1)
	actual2, _ := findLastNumber(line2)
	actual3, _ := findLastNumber(line3)
	actual4, _ := findLastNumber(line4)

	if actual1 != expected1 {
		t.Fatalf("Acutal 1 was: %s; Expected was %s", actual1, expected1)
	}

	if actual2 != expected2 {
		t.Fatalf("Acutal 2 was: %s; Expected was %s", actual2, expected2)
	}

	if actual3 != expected3 {
		t.Fatalf("Acutal 3 was: %s; Expected was %s", actual3, expected3)
	}

	if actual4 != expected4 {
		t.Fatalf("Acutal 4 was: %s; Expected was %s", actual4, expected4)
	}
}
