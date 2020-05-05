package offside

import (
	"fmt"
	"strings"
	"ysn/errors"
	"ysn/lexer"
	"ysn/token"
)

type offside struct {
	tokens []token.Token
	curr   int
}

func New(ls lexer.Lexerstream) offside {
	var tokens []token.Token
	mustIdent := false
	identLevel := 0
	identFound := 0
	line := 0
	for {
		tk, err := ls.Next()
		if err != nil {
			_, nomatch := err.(*errors.NoMatch)
			if nomatch {
				panic("error during lexing")
			}
			break
		}
		if tk.GetType() == "COLON" {
			mustIdent = true
		}
		if tk.GetType() == "LINESTART" {
			line++
			identFound = len(strings.ReplaceAll(tk.GetStr(), "\n", ""))
			if mustIdent {
				if identFound == (identLevel + 1) {
					mustIdent = false
					identLevel = identFound
					tokens = append(tokens, token.New("INDENT", ""))
				} else {
					panic(fmt.Sprintf("indentation error found line %v", line))
				}
			} else {
				if identFound > identLevel {
					panic(fmt.Sprintf("indentation error found line %v", line))
				} else if identFound < identLevel {
					for i := 0; i < (identLevel - identFound); i++ {
						tokens = append(tokens, token.New("DEDENT", ""))
					}
					identLevel = identFound
				} else {
					tokens = append(tokens, token.New("NEWLINE", "\n"))
				}
			}
		}
		if tk.GetType() != "LINESTART" {
			tokens = append(tokens, tk)
		}
	}
	if identLevel != 0 {
		for i := 0; i < identLevel; i++ {
			tokens = append(tokens, token.New("DEDENT", ""))
		}
	}
	if tokens[0].GetType() != "START" {
		panic("wrong start of source")
	}
	return offside{tokens, -1}
}

func (o *offside) Next() (tok token.Token, err error) {
	o.curr++
	if o.curr < len(o.tokens) {
		return o.tokens[o.curr], nil
	} else {
		return token.New(), &errors.EndOfSource{}
	}
}
