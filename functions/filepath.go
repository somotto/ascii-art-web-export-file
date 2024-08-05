package functions

import "path/filepath"

// FileName checks if the provided filename has a valid ".txt" extension.
func FileName(fileName string) string {
	if filepath.Ext(fileName) != ".txt" {
		return ""
	}
	switch fileName {
	case "standard.txt", "shadow.txt", "thinkertoy.txt":

		return fileName
	default:
		return ""
	}
}
