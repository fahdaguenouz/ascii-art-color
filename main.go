package main

import (
	
	"fmt"
	"os"
	"strings"
	"asciiartcolor/functions"
)

func main() {
	args := os.Args[1:]

	colors := map[string]string{
	"red" : "\033[31m",
	"green" : "\033[32m",
	"blue" : "\033[34m",
	"yellow" : "\033[33m",
	"cyan" : "\033[36m",
	"magenta" : "\033[35m",
	"white" : "\033[37m",
	"black" : "\033[30m",
	"gray" : "\033[90m",
	"purple" : "\033[35m",
	"orange" : "\033[38;5;214m" ,
	"reset" : "\033[0m" ,
	}


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
	// Check if the color exists in the map
	colorCode, exists := colors[colorname]
	if !exists {
		fmt.Println("Error: Invalid color name or format")
		return
	}
	
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

result := asciiartcolor.GenerateASCIIArt(input, asciiArt)
fmt.Print(colorCode + result + colors["reset"])

	
	}else if len(args) == 4{
		color := args[0]
		word:=args[1]
		input := args[2]
		banner := args[3]
		// Validate output option
		if !strings.HasPrefix(color, "--color=") {
			fmt.Println("Invalid option. Please use --color= < name of the color >")
			return
		}
	
		// Validate color name
		colorname := strings.TrimPrefix(color, "--color=")
		colorCode, exists := colors[colorname]
		if !exists {
			fmt.Println("Error: Invalid color name or format")
			return
		}
		
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
// Generate ASCII art for the input string and colorize the target word
result := asciiartcolor.GenerateASCIIArtLetter(input, asciiArt, word, colorCode, colors["reset"])

fmt.Print(result)

	}else {
		fmt.Println("Please enter this format: 'go run . [color] [a letter or word optionel ] [STRING] [BANNER]'")
		return
	}

	

}
