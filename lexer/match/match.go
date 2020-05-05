package match

import "fmt"

type Match struct {
	Start int
	End   int
}

func New(start int, end int) Match {
	return Match{start, end}
}

func (m *Match) Repr() string {
	return fmt.Sprintf("Match(Start=%v, End=%v)", m.Start, m.End)
}
