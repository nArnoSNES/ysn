package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"ysn/lexer/lexergenerator"
	"ysn/lexer/offside"
	"ysn/utils/helpers"
)

func main() {
	str, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	source := string(str)
	fmt.Println(source)
	source = helpers.Clean(source)
	fmt.Println(source)
	tokens := toklist()
	lg := lexergenerator.New()

	lg.Add("START", `---`, nil)
	lg.Add("COLON", `:`, nil)
	lg.Add("LINESTART", `\n(\t)*`, nil)

	iter := tokens.IterFunc()
	for kv, ok := iter(); ok; kv, ok = iter() {
		lg.Add(kv.Key.(string), kv.Value.(string), nil)
	}
	lg.Ignore(`[ \r\f\v]+`, nil)
	l := lg.Build()
	ls := l.Lex(source)
	o := offside.New(ls)
	for {
		tk, err := o.Next()
		if err != nil {
			break
		}
		fmt.Println(tk.Repr())
	}
}
