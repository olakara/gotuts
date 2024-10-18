// Package iterations can repeat string multiple times
package iterations

// Repeat function returns n number of repeats for the passed string
func Repeat(char string, count int) string {
	var result string
	for i := 0; i < count; i++ {
		result += char
	}
	return result
}
