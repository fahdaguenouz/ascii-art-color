package asciiartcolor

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to get the appropriate art file based on the banner type
func GetArtFile(banner string) (string, error) {
	standard := "./art/standard.txt"
	shadow := "./art/shadow.txt"
	thinkertoy := "./art/thinkertoy.txt"

	switch banner {
	case "standard":
		return standard, nil
	case "shadow":
		return shadow, nil
	case "thinkertoy":
		return thinkertoy, nil
	default:
		return "", fmt.Errorf("invalid banner. Please choose from: standard, shadow, thinkertoy")
	}
}

func ReadArtFile(art string) ([]string, error) {
	file, err := os.Open(art)
	if err != nil {
		return nil, fmt.Errorf("error opening the file: %v", err)
	}
	defer file.Close()

	var asciiArt []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		asciiArt = append(asciiArt, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading the file: %v", err)
	}

	return asciiArt, nil
}



// Function to generate ASCII art from the input string
func GenerateASCIIArt(input string, asciiArt []string) (string) {
	var result string
	lines := strings.Split(input, "\\n")

	for _, line := range lines {
		if line == "" {
			result += "\n"
			continue
		}
		for i := 1; i <= 8; i++ {
			for _, r := range line {
				if r < 32 || r > 126 {
					 fmt.Printf("invalid character: please enter a valid character between ASCII code 32 and 126")
					 return ""
				}
				index := 9*(int(r)-32) + i
				result += asciiArt[index]
			}
			result += "\n"
		}
	}

	return result
}
func GenerateASCIIArtLetter(input string, asciiArt []string, targetWord string, colorCode string, resetCode string) string {
	var result string
	lines := strings.Split(input, "\\n")

	// Process each line
	for _, line := range lines {
		if line == "" {
			result += "\n"
			continue
		}

		// Loop through each line of ASCII art (8 lines per character)
		for i := 1; i <= 8; i++ {
			for _, r := range line {
				if r < 32 || r > 126 {
					fmt.Printf("invalid character: please enter a valid character between ASCII code 32 and 126")
					return ""
				}
				index := 9*(int(r)-32) + i // Get the index of the ASCII art line for the character

				// Check if the current character matches part of the targetWord
				if strings.ContainsRune(targetWord, r) {
					// Apply the color to the matched character
					result += colorCode + asciiArt[index] + resetCode
				} else {
					// No match, just add the ASCII art without coloring
					result += asciiArt[index]
				}
			}
			result += "\n" // Add newline after each set of ASCII lines
		}
	}

	return result
}


