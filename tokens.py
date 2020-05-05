from collections import OrderedDict

tokens = OrderedDict() # Mandatory use of OrderedDict

# Inline values
tokens["INTEGER"] = "0x[0-9A-Fa-f]+|-?\d+"
tokens["STRING"] = '(""".*?""")|(".*?")|(\'.*?\')'
tokens["BOOLEAN"] = "True(?!\w)|False(?!\w)"

# Types
tokens["TYPE_INTEGER"] = "int(?!\w)"
tokens["TYPE_STRING"] = "str(?!\w)"
tokens["ENUM"] = "enum(?!\w)"

# Structurals
tokens["IF"] = "if(?!\w)"
tokens["ELIF"] = "elif(?!\w)"
tokens["ELSE"] = "else(?!\w)"

tokens["FOR"] = "for(?!\w)"
tokens["IN"] = "in(?!\w)"

tokens["WHILE"] = "while(?!\w)"
tokens["BREAK"] = "break(?!\w)"

tokens["SWITCH"] = "switch(?!\w)"
tokens["CASE"] = "case(?!\w)"
tokens["DEFAULT"] = "default(?!\w)"

tokens["MACRO"] = "macro(?!\w)"
tokens["RETURN"] = "return(?!\w)"

tokens["PASS"] = "pass(?!\w)"

# conditionnals
tokens["AND"] = "and(?!\w)"
tokens["OR"] = "or(?!\w)"
tokens["NOT"] = "not(?!\w)"

# object names
tokens["IDENTIFIER"] = "[a-zA-Z_][a-zA-Z0-9_]*"
tokens["DOT"] = "\."

# bin ops
tokens[">>"] = ">>"
tokens["<<"] = "<<"
tokens["&&"] = "&&"

# incr and decr
tokens["INCR"] = "\+\+"
tokens["DECR"] = "--"

# arithmetics
tokens["PLUS"] = "\+"
tokens["MINUS"] = "-"
tokens["MUL"] = "\*"
tokens["DIV"] = "/"
tokens["MOD"] = "%"

# comparaison
tokens["=="] = "=="
tokens["!="] = "!="
tokens[">="] = ">="
tokens["<="] = "<="
tokens[">"] = ">"
tokens["<"] = "<"

# enclosure
tokens["["] = "\["
tokens["]"] = "\]"
tokens["{"] = "\{"
tokens["}"] = "\}"
tokens["("] = "\("
tokens[")"] = "\)"
tokens[","] = ","

# specials
tokens["="] = "="
tokens["$"] = "\$"
tokens["&"] = "\&"
tokens["%"] = "%"

