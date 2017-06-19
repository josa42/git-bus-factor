package stringutils

import (
	"regexp"
	"strings"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// TrimLeadingTabs :
func TrimLeadingTabs(str string) string {
	str = RemoveSurroundingEmptyLines(str)
	lines := strings.Split(str, "\n")

	r, _ := regexp.Compile("^\\t+")
	for idx, line := range lines {
		lines[idx] = r.ReplaceAllString(line, "")
	}

	return strings.Join(lines, "\n")
}

// TrimPrefix :
func TrimPrefix(str string) string {

	str = RemoveSurroundingEmptyLines(str)
	lines := strings.Split(str, "\n")

	prefix := countLeadingSpaces(lines[0])

	for idx, line := range lines {

		removeEnd := min(prefix, len(line))

		lines[idx] = line[removeEnd:]
	}

	return strings.Join(lines, "\n")
}

// RemoveSurroundingEmptyLines :
func RemoveSurroundingEmptyLines(str string) string {
	lines := strings.Split(str, "\n")

	for len(lines) > 1 && strings.Trim(lines[0], "\t ") == "" {
		lines = append(lines[:0], lines[1:]...)
	}

	for len(lines) > 1 && strings.Trim(lines[len(lines)-1], "\t ") == "" {
		length := len(lines)
		lines = lines[:length-1]
	}

	return strings.Join(lines, "\n")
}

func countLeadingSpaces(line string) int {
	count := 0
	for _, rune := range line {
		if rune != '\t' {
			break
		}

		count++
	}
	return count
}
