// +build ignore

// This program generates all-colors.go, which allows us to compile
// our all-colors.json dataset inside of our Go binary, for easy access.
//
// This file is designed to be only be invoked with go:generate
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	prefix := []byte("package main \n\nvar allColorsJSONString = `\n")
	suffix := []byte("`")

	// The color names in all.json are sourced from Pantone, HTML, X11 and NTC. For the original source, see
	// https://github.com/colorjs/color-namer/tree/abb3d184ab63db9327908319cc45b55c91493bb7/lib/colors
	allColorsJSONBytes, err := ioutil.ReadFile("all-colors.json")

	if err != nil {
		fmt.Println("error:", err)
	}

	file, err := os.Create("all-colors.go")

	if err != nil {
		fmt.Println("error:", err)
	}

	file.Write(prefix)
	file.Write(allColorsJSONBytes)
	file.Write(suffix)
}
