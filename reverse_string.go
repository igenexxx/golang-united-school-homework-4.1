package reverse_string

import "fmt"

func reverse(input string) (reversed string) {
	for _, runeUnit := range input {
		reversed = fmt.Sprintf("%c%s", runeUnit, reversed)
	}

	return
}

func ReverseString(input string) (output string) {
	output = reverse(input)

	return
}
