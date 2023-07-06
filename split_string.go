package gotools

// Splits a string into a string array and returns it
func SplitString(target string) []string {
	to_return := []string{}
	chars := []rune(target)
	for i := 0; i < len(chars); i++ {
		to_return = append(to_return, string(chars[i]))
	}
	return to_return
}
