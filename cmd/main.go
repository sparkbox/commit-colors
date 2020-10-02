package main

import (
	"fmt"
	"github.com/gookit/color"
	"os"
)

//go:generate go run gen.go

func main() {
	myHexID, err := GetHexIDFromCommitID(GetCommitIDFromArgs(os.Args))

	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	// Get all the relevant color data
	myColorSwatchBackgroundHedID := myHexID
	myColorSwatchForegroundHexID := "FFFFFF"
	myColorSwatchStyle := color.HEXStyle(myColorSwatchForegroundHexID, myColorSwatchBackgroundHedID)
	myColorName := GetColorName(myHexID)

	// Print the color swatch to the terminal
	fmt.Println("")
	myColorSwatchStyle.Println("       ")
	myColorSwatchStyle.Printf("       ")
	fmt.Printf(" Your commit color is %s\n", myColorName)
	myColorSwatchStyle.Println("       ")
	fmt.Printf("#%s\n", myHexID)
	fmt.Println("")
}
