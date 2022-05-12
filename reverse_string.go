package reverse_string

func reverse(input string) string {
	var result []uint8
	length := len(input) - 1

	for i, _ := range input {
		result = append(result, input[length-i])
	}

	return string(result)
}

func ReverseString(input string) (output string) {
	// solution goes here
	output = reverse(input)

	return
}
