package gotools

import "strings"

// Wrapper for the strings.HasPrefix function, but takes in a array contaning strings to look for
// and returns a map in the format of { "<key>":"<true/false>" }
func StartsWith(target string, selection []string) map[string]bool {
	to_return := make(map[string]bool)

	for _, value := range selection {
		to_return[value] = strings.HasPrefix(target, value)
	}

	return to_return
}
