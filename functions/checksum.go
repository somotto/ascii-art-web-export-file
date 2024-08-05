package functions

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

// expectedHashes defines the expected SHA-256 hashes for specific files.
var expectedHashes = map[string]string{
	"standard.txt":   "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf",
	"shadow.txt":     "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73",
	"thinkertoy.txt": "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3",
}

// VerifyFileChecksum verifies the SHA-256 checksum of a file against an expected value.
func VerifyFileChecksum(banner, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		// Handle the error if the file does not exist.
		if os.IsNotExist(err) {
			return ErrFileMissing
		}
		return err
	}
	defer file.Close()

	// Create a new SHA-256 hash object.
	hash := sha256.New()

	// Copy the file contents into the hash object.
	if _, err := io.Copy(hash, file); err != nil {
		return ErrFileCorrupted
	}

	// Calculate the hexadecimal representation of the hash.
	calculatedHash := fmt.Sprintf("%x", hash.Sum(nil))

	// If no banner is provided, use the filename as the banner.
	if banner == "" {
		banner = filename
	}

	// Retrieve the expected hash value from the map.
	expectedHash, ok := expectedHashes[banner]
	if !ok {
		return ErrFileCorrupted
	}

	// Compare the calculated hash with the expected hash.
	if calculatedHash != expectedHash {
		return ErrFileCorrupted
	}

	return nil
}
