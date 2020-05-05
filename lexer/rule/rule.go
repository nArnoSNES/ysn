package rule

import (
	"fmt"
	"regexp"
	"ysn/lexer/match"
)

type Rule struct {
	name    string
	pattern string
	flags   string
	re      regexp.Regexp
}

func New(name string, pattern string, flags interface{}) Rule {
	tflags, okflag := flags.(string)
	if okflag {
		return Rule{name, pattern, tflags, *regexp.MustCompile(fmt.Sprintf("(?%s)%s", tflags, pattern))}
	} else {
		return Rule{name, pattern, "", *regexp.MustCompile(pattern)}
	}
}

func (r *Rule) Name() string {
	return r.name
}

func findStringIndexPos(r regexp.Regexp, s string, p int) []int {
	m := r.FindStringIndex(s[p:])
	if m != nil && m[0] == 0 {
		m[0] = p
		m[1] = m[1] + p
		return m
	} else {
		return nil
	}
}

func (r *Rule) Matches(s string, p int) match.Match {
	m := findStringIndexPos(r.re, s, p)
	if m != nil {
		return match.New(m[0], m[1])
	} else {
		return match.New(-1, -1)
	}
}
