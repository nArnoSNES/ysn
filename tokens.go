package main

import "github.com/cevaris/ordered_map"

func toklist() *ordered_map.OrderedMap {

	tokens := ordered_map.NewOrderedMap()

	// Inline values
	tokens.Set("INTEGER", `0x[0-9A-Fa-f]+|-?\d+`)
	tokens.Set("STRING", `(""".*?""")|(".*?")|('.*?')`)
	tokens.Set("BOOLEAN", `True(?!\w)|False(?!\w)`)

	// Types
	tokens.Set("TYPE_INTEGER", `int(?!\w)`)
	tokens.Set("TYPE_STRING", `str(?!\w)`)
	tokens.Set("ENUM", `enum(?!\w)`)

	// Structurals
	tokens.Set("IF", `if(?!\w)`)
	tokens.Set("ELIF", `elif(?!\w)`)
	tokens.Set("ELSE", `else(?!\w)`)

	tokens.Set("FOR", `for(?!\w)`)
	tokens.Set("IN", `in(?!\w)`)

	tokens.Set("WHILE", `while(?!\w)`)
	tokens.Set("BREAK", `break(?!\w)`)

	tokens.Set("SWITCH", `switch(?!\w)`)
	tokens.Set("CASE", `case(?!\w)`)
	tokens.Set("DEFAULT", `default(?!\w)`)

	tokens.Set("MACRO", `macro(?!\w)`)
	tokens.Set("RETURN", `return(?!\w)`)

	tokens.Set("PASS", `pass(?!\w)`)

	// conditionnals
	tokens.Set("AND", `and(?!\w)`)
	tokens.Set("OR", `or(?!\w)`)
	tokens.Set("NOT", `not(?!\w)`)

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
