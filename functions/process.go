package functions

import (
	"errors"
	"os"
	"strings"
)

var (
	ErrFileCorrupted = errors.New("file is corrupted")
	ErrFileMissing   = errors.New("file is missing")
)

// ReadFromFile gets the filename and splits it into an array of strings.
func Readfile(fileName string) ([]string, error) {
	err := VerifyFileChecksum("", fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, ErrFileMissing
		}
		return nil, ErrFileCorrupted
	}

	file, err := os.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, ErrFileMissing
		}
		return nil, err
	}
	if string(file) == "" {
		return nil, ErrFileCorrupted
	}
	var lines []string
	if fileName == "thinkertoy.txt" {
		lines = strings.Split(string(file), "\r\n")
	} else {
		lines = strings.Split(string(file), "\n")
	}
	return lines, nil
}
