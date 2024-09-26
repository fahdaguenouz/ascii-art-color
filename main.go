package main

import (
	
	"fmt"
	"os"
	"strings"
	"asciiartcolor/functions"
)

func main() {
	args := os.Args[1:]

	// Predefined list of named colors (CSS-like)


	if len(args) == 3 {
	color := args[0]
	input := args[1]
	banner := args[2]
	// Validate output option
	if !strings.HasPrefix(color, "--color=") {
		fmt.Println("Invalid option. Please use --color= < name of the color >")
		return
	}

	// Validate color name
	colorname := strings.TrimPrefix(color, "--color=")
	// if ValidateColor(colorname)==false {
	// 	fmt.Println("Error: Invalid color name or format")
	// 	return 
	// } 
	fmt.Println(colorname)
// Get the appropriate ASCII art file
		artFile, err := asciiartcolor.GetArtFile(banner)
if err != nil {
	fmt.Println(err)
	return
}
// Read the ASCII art file
asciiArt, err := asciiartcolor.ReadArtFile(artFile)
if err != nil {
	fmt.Println(err)
	return
}


// Generate ASCII art for the input string

fmt.Print(asciiartcolor.GenerateASCIIArt(input, asciiArt))

	
	}else if len(args) == 4{
		color := args[0]
		letter:=args[1]
		input := args[2]
		banner := args[3]
		// Validate output option
		if !strings.HasPrefix(color, "--color=") {
			fmt.Println("Invalid option. Please use --color= < name of the color >")
			return
		}
	
		// Validate color name
		colorname := strings.TrimPrefix(color, "--color=")
		// if ValidateColor(colorname)==false {
		// 	fmt.Println("Error: Invalid color name or format")
		// 	return 
		// } 
		fmt.Println(colorname,letter)
	// Get the appropriate ASCII art file
			artFile, err := asciiartcolor.GetArtFile(banner)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Read the ASCII art file
	asciiArt, err := asciiartcolor.ReadArtFile(artFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	
	fmt.Print(asciiartcolor.GenerateASCIIArt(input, asciiArt))

	}else {
		fmt.Println("Please enter this format: 'go run . [color] [a letter or word optionel ] [STRING] [BANNER]'")
		return
	}

	

}