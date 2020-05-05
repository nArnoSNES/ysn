package token

import (
	"fmt"
	"rply/token/sourceposition"
	"rply/utils/is"
)

type token struct {
	name      string
	value     string
	sourcepos sourceposition.Sourceposition
}

func (tk *token) Repr() string {
	return fmt.Sprintf("token(name=%v, value=%v)", tk.name, tk.value)
}

func (tk *token) Equal(other token) bool {
	return (tk.name == other.name && tk.value == other.value)
}

func (tk *token) GetType() string {
	return tk.name
}

func (tk *token) GetStr() string {
	return tk.value
}

func (tk *token) GetPos() sourceposition.Sourceposition {
	return tk.sourcepos
}

func New(args ...interface{}) token {
	switch len(args) {
	case 3:
		_, isSourcepos := args[2].(sourceposition.Sourceposition)
		if is.String(args[0]) && is.String(args[1]) && isSourcepos {
			return token{args[0].(string), args[1].(string), args[2].(sourceposition.Sourceposition)}
		}
		fallthrough
	case 2:
		if is.String(args[0]) && is.String(args[1]) {
			return token{args[0].(string), args[1].(string), sourceposition.New()}
		}
		fallthrough
	default:
		return token{"", "", sourceposition.New()}
	}
}
