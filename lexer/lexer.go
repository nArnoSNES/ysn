package lexer

import (
	"strings"
	"ysn/errors"
	"ysn/lexer/match"
	"ysn/lexer/rule"
	"ysn/token"
	"ysn/token/sourceposition"
)

type Lexer struct {
	rules   []rule.Rule
	ignores []rule.Rule
}

func New(rules []rule.Rule, ignores []rule.Rule) Lexer {
	return Lexer{rules, ignores}
}

func (l *Lexer) Lex(s string) Lexerstream {
	return Lexerstream{*l, s, 0, 1, 1}
}

type Lexerstream struct {
	lexer  Lexer
	source string
	idx    int
	lineno int
	colno  int
}

func (ls *Lexerstream) updatePos(m match.Match) int {
	ls.idx = m.End
	ls.lineno += strings.Count(ls.source[m.Start:m.End], "\n")
	lastnl := strings.LastIndex(ls.source[0:m.Start], "\n")
	if lastnl == -1 {
		return m.Start + 1
	} else {
		return m.Start - lastnl
	}
}

func (ls *Lexerstream) Next() (tok token.Token, err error) {
	var matchFound bool
	for {
		if ls.idx >= len(ls.source) {
			return token.New(), &errors.EndOfSource{}
		}
		matchFound = false
		for _, r := range ls.lexer.ignores {
			m := r.Matches(ls.source, ls.idx)
			matchFound = (m.Start != -1)
			if matchFound {
				ls.colno = ls.updatePos(m)
				break
			}
		}
		if !matchFound {
			break
		}
	}
	for _, r := range ls.lexer.rules {
		m := r.Matches(ls.source, ls.idx)
		if m.Start != -1 {
			lineno := ls.lineno
			ls.colno = ls.updatePos(m)
			sp := sourceposition.New(m.Start, lineno, ls.colno)
			tk := token.New(r.Name(), ls.source[m.Start:m.End], sp)
			return tk, nil
		}
	}
	return token.New(), &errors.NoMatch{}
}
