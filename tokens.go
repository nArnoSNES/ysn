package main

import "github.com/cevaris/ordered_map"

func toklist() *ordered_map.OrderedMap {

	tokens := ordered_map.NewOrderedMap()

	// Inline values
	tokens.Set("INTEGER", `0x[0-9A-Fa-f]+|-?\d+`)
	tokens.Set("STRING", `(""".*?""")|(".*?")|('.*?')`)
	tokens.Set("BOOLEAN", `True|False`)

	// Types
	tokens.Set("TYPE_INTEGER", `int`)
	tokens.Set("TYPE_STRING", `str`)
	tokens.Set("ENUM", `enum`)

	// Structurals
	tokens.Set("IF", `if`)
	tokens.Set("ELIF", `elif`)
	tokens.Set("ELSE", `else`)

	tokens.Set("FOR", `for`)
	tokens.Set("IN", `in`)

	tokens.Set("WHILE", `while`)
	tokens.Set("BREAK", `break`)

	tokens.Set("SWITCH", `switch`)
	tokens.Set("CASE", `case`)
	tokens.Set("DEFAULT", `default`)

	tokens.Set("MACRO", `macro`)
	tokens.Set("RETURN", `return`)

	tokens.Set("PASS", `pass`)

	// conditionnals
	tokens.Set("AND", `and`)
	tokens.Set("OR", `or`)
	tokens.Set("NOT", `not`)

	// object names
	tokens.Set("IDENTIFIER", `[a-zA-Z_][a-zA-Z0-9_]*`)
	tokens.Set("DOT", `\.`)

	// bin ops
	tokens.Set(">>", `>>`)
	tokens.Set("<<", `<<`)
	tokens.Set("&&", `\&\&`)

	// incr and decr
	tokens.Set("INCR", `\+\+`)
	tokens.Set("DECR", `--`)

	// arithmetics
	tokens.Set("PLUS", `\+`)
	tokens.Set("MINUS", `-`)
	tokens.Set("MUL", `\*`)
	tokens.Set("DIV", `/`)
	tokens.Set("MOD", `%`)

	// comparaison
	tokens.Set("==", `==`)
	tokens.Set("!=", `!=`)
	tokens.Set(">=", `>=`)
	tokens.Set("<=", `<=`)
	tokens.Set(">", `>`)
	tokens.Set("<", `<`)

	// enclosure
	tokens.Set("[", `\[`)
	tokens.Set("]", `\]`)
	tokens.Set("{", `\{`)
	tokens.Set("}", `\}`)
	tokens.Set("(", `\(`)
	tokens.Set(")", `\)`)
	tokens.Set(",", `,`)

	// specials
	tokens.Set("=", `=`)
	tokens.Set("$", `\$`)
	tokens.Set("&", `\&`)
	tokens.Set("%", `%`)

	return tokens
}
