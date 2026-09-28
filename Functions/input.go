package Functions

// isValidInput reports whether input contains only printable ASCII characters.
func IsValidInput(input string) bool {
	for _, character := range input {
		if character < 32 || character > 126 {
			return false
		}
	}
	return true
}

// IsEmpty reports whether every string in arr is empty.
func IsEmpty(arr []string) bool {
	for _, value := range arr {
		if value != "" {
			return false
		}
	}
	return true
}
