package main

import (
	"testing"
)

func TestFirstNumber(t *testing.T) {
	line1 := "dssmtmrkonedbbhdhjbf9hq"
	line2 := "bsxxzhkmmfcslmdhhrgf9seven665lsknmbpgj"
	line3 := "9"

	expected := "9"

	actual1, _ := findFirstNumber(line1)
	actual2, _ := findFirstNumber(line2)
	actual3, _ := findFirstNumber(line3)

	if actual1 != expected {
		t.Fatalf("Acutal 1 was: %s; Expected was %s", actual1, expected)
	}

	if actual2 != expected {
		t.Fatalf("Acutal 1 was: %s; Expected was %s", actual2, expected)
	}

	if actual3 != expected {
		t.Fatalf("Acutal 1 was: %s; Expected was %s", actual3, expected)
	}
}

func TestLastNumber(t *testing.T) {
	line1 := "dssmtmrkonedbbhdhjbf9hq"
	line2 := "bsxxzhkmmfcslmdhhrgf9seven669lsknmbpgj"
	line3 := "9"

	expected := "9"

	actual1, _ := findLastNumber(line1)
	actual2, _ := findLastNumber(line2)
	actual3, _ := findLastNumber(line3)

	if actual1 != expected {
		t.Fatalf("Acutal 1 was: %s; Expected was %s", actual1, expected)
	}

	if actual2 != expected {
		t.Fatalf("Acutal 1 was: %s; Expected was %s", actual2, expected)
	}

	if actual3 != expected {
		t.Fatalf("Acutal 1 was: %s; Expected was %s", actual3, expected)
	}
}
