package token

import (
	"fmt"
	"ysn/token/sourceposition"
	"ysn/utils/is"
)

type Token struct {
	name      string
	value     string
	sourcepos sourceposition.Sourceposition
}

func (tk *Token) Repr() string {
	return fmt.Sprintf("Token(name='%v', value='%v')", tk.name, tk.value)
}

func (tk *Token) Equal(other Token) bool {
	return (tk.name == other.name && tk.value == other.value)
}

func (tk *Token) GetType() string {
	return tk.name
}

func (tk *Token) GetStr() string {
	return tk.value
}

func (tk *Token) GetPos() sourceposition.Sourceposition {
	return tk.sourcepos
}

func New(args ...interface{}) Token {
	switch len(args) {
	case 3:
		_, isSourcepos := args[2].(sourceposition.Sourceposition)
		if is.String(args[0]) && is.String(args[1]) && isSourcepos {
			return Token{args[0].(string), args[1].(string), args[2].(sourceposition.Sourceposition)}
		}
		fallthrough
	case 2:
		if is.String(args[0]) && is.String(args[1]) {
			return Token{args[0].(string), args[1].(string), sourceposition.New()}
		}
		fallthrough
	default:
		return Token{"", "", sourceposition.New()}
	}
}
