package lexergenerator

import (
	"ysn/lexer"
	"ysn/lexer/rule"
)

type lexergenerator struct {
	rules   []rule.Rule
	ignores []rule.Rule
}

func New() lexergenerator {
	return lexergenerator{}
}

func (lg *lexergenerator) Add(name string, pattern string, flags interface{}) {
	lg.rules = append(lg.rules, rule.New(name, pattern, flags))
}

func (lg *lexergenerator) Ignore(pattern string, flags interface{}) {
	lg.ignores = append(lg.ignores, rule.New("", pattern, flags))
}

func (lg *lexergenerator) Build() lexer.Lexer {
	return lexer.New(lg.rules, lg.ignores)
}
