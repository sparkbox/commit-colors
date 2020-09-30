package main

import "testing"

func TestCommitIDTruncation(t *testing.T) {
	longCommitID := "ABC123DEF456"
	expectedTruncatedID := "ABC123"

	truncatedID, err := GetHexIDFromCommitID(longCommitID)

	if truncatedID != expectedTruncatedID || err != nil {
		t.Errorf("Truncation was incorrect, got: %s, expected: %s.", truncatedID, expectedTruncatedID)
	}
}

func TestInvalidCommitID(t *testing.T) {
	nonHexadecimalID := "PACD21"
	_, err := GetHexIDFromCommitID(nonHexadecimalID)

	if err == nil {
		t.Errorf("Invalid commit '%s' wasn't caught", nonHexadecimalID)
	}
}

func TestShortCommitID(t *testing.T) {
	shortID := "ABC"
	_, err := GetHexIDFromCommitID(shortID)

	if err == nil {
		t.Errorf("Invalid commit '%s' wasn't caught", shortID)
	}
}

func TestGetColorNameExactMatch(t *testing.T) {
	expectedColorName := "Yellow"
	colorName := GetColorName("FFFF00")

	if colorName != expectedColorName {
		t.Errorf("Color Name was incorrect, got: %s, expected: %s.", colorName, expectedColorName)
	}
}

func TestGetColorNameNearMatch(t *testing.T) {
	expectedColorName := "Pomegranate"
	colorName := GetColorName("DF2301")

	if colorName != expectedColorName {
		t.Errorf("Color Name was incorrect, got: %s, expected: %s.", colorName, expectedColorName)
	}
}
