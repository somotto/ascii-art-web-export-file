package functions

import (
	"strings"
)

// AsciiArt generates ASCII art for the given words.
func AsciiArt(input string, inputFile []string) string {
    var result strings.Builder
    lines := strings.Split(input, "\r\n")
    
    for _, line := range lines {
        if line == "" {
            result.WriteString(strings.Repeat("\n", 8))
        } else {
            for i := 0; i < 8; i++ {
                for _, char := range line {
                    if char < 32 || char > 126 {
                        continue // Skip non-printable characters
                    }
                    start := (int(char-32) * 9) + 1 + i
                    if start < len(inputFile) {
                        result.WriteString(inputFile[start])
                    }
                }
                result.WriteString("\n")
            }
        }
    }
    
    return result.String()
}
