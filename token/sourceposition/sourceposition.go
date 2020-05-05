package sourceposition

import (
	"fmt"
	"rply/utils/is"
)

type Sourceposition struct {
	Idx    int
	Lineno int
	Colno  int
}

func (sp *Sourceposition) Repr() string {
	return fmt.Sprintf("Sourceposition(Idx=%v, Lineno=%v, Colno=%v)", sp.Idx, sp.Lineno, sp.Colno)
}

func New(args ...interface{}) Sourceposition {
	if len(args) == 3 && is.Int(args[0]) && is.Int(args[1]) && is.Int(args[2]) {
		return Sourceposition{args[0].(int), args[1].(int), args[2].(int)}
	} else {
		return Sourceposition{0, 0, 0}
	}
}
