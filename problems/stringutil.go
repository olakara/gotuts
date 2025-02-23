package main

func reverse(inputString string) string {

	length := len(inputString)
	outputString := ""

	for i := length - 1; i >= 0; i-- {
		outputString += string(inputString[i])
	}

	return outputString
}
