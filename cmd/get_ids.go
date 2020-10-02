package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
)

// GetCommitIDFromArgs returns the first CLI argument and exits gracefully if it's missing
func GetCommitIDFromArgs(args []string) string {
	// Note: os.Args[] always contains at least one value (the path to the program).
	// Any additional values that exist are the arguments passed to the program.
	if len(args) == 1 {
		fmt.Println("commit-colors expects you to provide your commit ID as an argument")
		os.Exit(3)
	}

	return args[1]
}

// GetHexIDFromCommitID checks an input string and truncates, if necessary, to return a valid HexID
func GetHexIDFromCommitID(commitID string) (string, error) {
	isValidHexadecimal, err := regexp.MatchString("^[0-9a-fA-F]+$", commitID)

	if err != nil {
		fmt.Println("error:", err)
	}

	if !isValidHexadecimal || len(commitID) < 6 {
		return "", errors.New("The provided commit ID was invalid")
	}

	return commitID[0:6], nil
}
