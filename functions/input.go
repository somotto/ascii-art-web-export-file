package functions

// InputString checks if the input string contains only valid ASCII characters.
// including newline and carriage return.
func InputString(input string) bool {
	for _, char := range input {
		if char != '\n' && char != '\r' && (char < 32 || char > 126) {
			return false
		}
	}
	return true
}
