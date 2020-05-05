package main

import (
	"fmt"
	"rply/lexer/rule"
	"rply/token"
	"rply/token/sourceposition"
)

func main() {
	fmt.Println("Test sourceposition")
	sp := sourceposition.New()
	sp.Idx = 4
	fmt.Println(sp.Repr())
	fmt.Println("Test token")
	tk := token.New("TITI", "toto")
	fmt.Println(tk.Repr())
	tk2 := token.New("TITI", "toto", sp)
	fmt.Println("Test token equality")
	fmt.Println(tk.Equal(tk2))
	fmt.Println("Test token pos access")
	sp2 := tk2.GetPos()
	fmt.Println(sp2.Repr())
	r := rule.New("test", `\+ `, nil)
	m := r.Matches("2 + 3", 2)
	if m.Start == -1 {
		fmt.Println("not found")
	} else {
		fmt.Println(m.Repr())
	}
}
