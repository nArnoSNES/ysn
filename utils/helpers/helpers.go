package helpers

import (
	"regexp"
	"strings"
)

func removeComments(source string) string {
	comments := regexp.MustCompile(`(#.*)(?:\n|\z)`)

	for comment := comments.FindStringSubmatchIndex(source); comment != nil; comment = comments.FindStringSubmatchIndex(source) {
		source = source[0:comment[2]] + source[comment[3]:]
	}

	return source
}

func removeEmptyLines(source string) string {
	multiline := regexp.MustCompile(`([\s]+)(?:\n)`)

	for line := multiline.FindStringSubmatchIndex(source); line != nil; line = multiline.FindStringSubmatchIndex(source) {
		source = source[0:line[2]] + source[line[3]:]
	}

	return source
}

func fixIdents(source string) string {
	var identString string
	idents := regexp.MustCompile(`^[\s]+`)
	notspaces := regexp.MustCompile(`[^ ]`)
	splitSource := strings.Split(source, "\n")

	for _, line := range splitSource {
		m := idents.FindStringIndex(line)
		if m != nil {
			identString = strings.Repeat(" ", m[1])
			break
		}
	}

	var newSource []string
	for _, line := range splitSource {
		m := notspaces.FindStringIndex(line)
		if m != nil {
			newSource = append(newSource, strings.ReplaceAll(line[0:m[0]], identString, "\t")+line[m[0]:])
		} else {
			newSource = append(newSource, line)
		}
	}

	return strings.Join(newSource, "\n")
}

func Clean(source string) string {
	return fixIdents(removeEmptyLines(removeComments(source)))
}
