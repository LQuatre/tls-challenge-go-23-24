package piscine

func LoafOfBread(str string) string {
	result := ""
	var count int
	// loafofbread
	// Write a function LoafOfBread() that takes a string and returns another one with words of 5 characters and skips the next character followed by newline \n.

	// If there is a space in the middle of a string it should ignore it and get the next character until getting to a length of 5.
	// If the string is less than 5 characters return "Invalid Output\n".

	if len(str) < 5 {
		return "Invalid Output\n"
	}
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			continue
		}
		if count == 5 {
			result += " "
			count = -1
		} else {
			result += string(str[i])
		}
		count++
	}
	return result + "\n"
}
