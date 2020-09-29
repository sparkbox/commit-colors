package main

import (
	"fmt"
	"github.com/gookit/color"
	"os"
	"regexp"
)

//go:generate go run gen.go

func main() {
	myHexID := getHexIDfromCommitID(getFirstArg())
	myColorSwatchBackgroundHedID := myHexID
	myColorSwatchForegroundHexID := "FFFFFF"
	myColorSwatchStyle := color.HEXStyle(myColorSwatchForegroundHexID, myColorSwatchBackgroundHedID)
	myColorName := GetColorName(myHexID)

	fmt.Println("")
	myColorSwatchStyle.Println("       ")
	myColorSwatchStyle.Printf("       ")
	fmt.Printf(" Your commit color is %s\n", myColorName)
	myColorSwatchStyle.Println("       ")
	fmt.Printf("#%s\n", myHexID)
	fmt.Println("")
}

func getFirstArg() string {
	// Note: os.Args[] always contains at least one value (the path to the program).
	// Any additional values that exist are the arguments passed to the program.
	if len(os.Args) == 1 {
		fmt.Println("commit-colors expects you to provide your commit ID as an argument")
		os.Exit(3)
	}

	return os.Args[1]
}

func getHexIDfromCommitID(commitID string) string {
	isValidHexadecimal, err := regexp.MatchString("^[0-9a-fA-F]+$", commitID)

	if err != nil {
		fmt.Println("error:", err)
	}

	if !isValidHexadecimal || len(commitID) < 6 {
		fmt.Println("The provided commit ID was invalid")
		os.Exit(3)
	}

	return commitID[0:6]
}
