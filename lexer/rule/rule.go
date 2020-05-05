package rule

import (
	"fmt"
	"regexp"
	"rply/lexer/match"
)

type rule struct {
	name    string
	pattern string
	flags   string
	re      regexp.Regexp
}

func New(name string, pattern string, flags interface{}) rule {
	tflags, okflag := flags.(string)
	if okflag {
		return rule{name, pattern, tflags, *regexp.MustCompile(fmt.Sprintf("(?%s)%s", tflags, pattern))}
	} else {
		return rule{name, pattern, "", *regexp.MustCompile(pattern)}
	}
}

func findStringIndexPos(r regexp.Regexp, s string, p int) []int {
	l := r.FindStringIndex(s)
	if l != nil && l[0] == p {
		return l
	} else {
		return nil
	}
}

func (r *rule) Matches(s string, p int) match.Match {
	m := findStringIndexPos(r.re, s, p)
	if m != nil {
		return match.New(m[0], m[1])
	} else {
		return match.New(-1, -1)
	}
}
