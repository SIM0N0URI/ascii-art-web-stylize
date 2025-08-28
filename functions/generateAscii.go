package asciiart

import (
	"fmt"
	"os"
	"strings"
)

// GenerateASCII converts input text into ASCII art using the specified banner font.
// It reads the font file, validates input, and builds the formatted output.
func GenerateASCII(input, banner string) (string, error) {
	bannerFile := "banners/" + banner + ".txt"

	input = strings.ReplaceAll(input, "\r\n", "\n")

	if len(input) > 1000 {
		return "", fmt.Errorf("max-length exceeded")
	}

	fontData, err := os.ReadFile(bannerFile)
	if err != nil {
		return "", fmt.Errorf("failed to read banner file")
	}
	if input == "" {
		return "", fmt.Errorf("empty string")
	}
	banners := strings.Split(strings.ReplaceAll(string(fontData), "\r\n", "\n"), "\n")

	textLines := strings.Split(input, "\n")

	result := ""

	for _, line := range textLines {
		if line == "" {
			result += "\n"
			continue
		}
		for i := 0; i <= 8; i++ {
			for _, char := range line {
				if char < ' ' || char > '~' {
					return "", fmt.Errorf("unsupported character(s)")
				}
				lineindex := int(char-32)*9 + i
				result += banners[lineindex]
			}
			result += "\n"
		}
	}
	return result, nil
}
