package functions

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestGraphic(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		expected []string
		banner   string
	}{
		{"Single letter", "H", artH, "standard"},
		{"Hello art", "Hello", artHello, "standard"},
		{"Single word", "K", arth, "shadow"},
		{"hello art", "hello", arthello, "shadow"},
		{"Letter", "X", artth, "thinkertoy"},
		{"Word", "Xoxo", artXoxo, "thinkertoy"},
	}

	for _, test := range tests {
		file, err := os.ReadFile("../" + test.banner + ".txt")
		if err != nil {
			fmt.Println("Error reading file", err)
			t.FailNow()
		}

		asciiArtString := string(file)

		var lines []string
		if test.banner == "thinkertoy" {
			lines = strings.Split(asciiArtString, "\r\n")
		} else {
			lines = strings.Split(asciiArtString, "\n")
		}
		t.Run(test.name, func(t *testing.T) {
			output := AsciiArt(test.str, lines)
			expected := artString(test.expected)
			if output != expected {
				t.Errorf("got: \n%v\nexpected: \n%v\n", output, expected)
			}
		})
	}
}

func artString(art []string) string {
	return strings.Join(art, "\n") + "\n"
}

var artH = []string{
	" _    _  ",
	"| |  | | ",
	"| |__| | ",
	"|  __  | ",
	"| |  | | ",
	"|_|  |_| ",
	"         ",
	"         ",
}

var artHello = []string{
	" _    _          _   _          ",
	"| |  | |        | | | |         ",
	"| |__| |   ___  | | | |   ___   ",
	`|  __  |  / _ \ | | | |  / _ \  `,
	"| |  | | |  __/ | | | | | (_) | ",
	`|_|  |_|  \___| |_| |_|  \___/  `,
	"                                ",
	"                                ",
}

var arth = []string{
	"         ",
	"_|    _| ",
	"_|  _|   ",
	"_|_|     ",
	"_|  _|   ",
	"_|    _| ",
	"         ",
	"         ",
}

var arthello = []string{
	"                                 ",
	"_|                _| _|          ",
	"_|_|_|     _|_|   _| _|   _|_|   ",
	"_|    _| _|_|_|_| _| _| _|    _| ",
	"_|    _| _|       _| _| _|    _| ",
	"_|    _|   _|_|_| _| _|   _|_|   ",
	"                                 ",
	"                                 ",
}

var artth = []string{
	"      ",
	"o   o ",
	` \ /  `,
	"  O   ",
	` / \  `,
	"o   o ",
	"      ",
	"      ",
}

var artXoxo = []string{
	"                  ",
	"o   o             ",
	` \ /              `,
	`  O   o-o \ / o-o `,
	` / \  | |  o  | | `,
	`o   o o-o / \ o-o `,
	"                  ",
	"                  ",
}
