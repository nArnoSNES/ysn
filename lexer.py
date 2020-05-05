from rply import LexerGenerator, Token
from tokens import tokens
from helpers import remove_comments, remove_empty_lines, fix_idents
import sys

lg = LexerGenerator()

# Mandatory token used by lexer for off-side rule.
lg.add("START", "---")
lg.add("COLON",":")
# Line start are actually a newline plus following tabs if exists to find indentation.
# To make this way of parsing line possible, code must start with a "START" token (really, any not significant string would do).
lg.add("LINESTART", "\n(\t)*")

# feed in generator a set of token names and regexes they match
for k in tokens.keys():
    lg.add(k,tokens[k])

# ignore whitespace
lg.ignore("[ \r\f\v]+")

lexer = lg.build()

def off_side(token_stream):

    # Search for COLON and LINESTART in Token Stream to insert INDENT/DEDENT Tokens
    # Very crude implementation of off-side rule (mandatory indent).
    # If a COLON is found, must_ident become true
    # If must_ident, next LINESTART must contains current ident_level + 1 (to INDENT) or raise an error
    # If not must_ident, when a LINESTART is found, if the ident found is lower than current ident_level, DEDENT
    l = []
    ident_level = 0
    must_ident = False
    line = 0
    for t in token_stream:
        if t.gettokentype() == "COLON":
            must_ident = True
        if t.gettokentype() == "LINESTART":
            line += 1
            ident_found = len(t.value.replace("\n", ""))
            if must_ident:
                if ident_found == ident_level + 1:
                    must_ident = False
                    ident_level = ident_found
                    l.append(Token("INDENT", ""))
                else:
                    raise Exception("Indent error line %i" % line)
            else:
                if ident_found > ident_level:
                    raise Exception("Indent error line %i" % line)
                elif ident_found < ident_level:
                    l.extend([Token("DEDENT", "")] * (ident_level - ident_found))
                    ident_level = ident_found
                else:
                    l.append(Token("NEWLINE", "\n"))
        if t.gettokentype() != "LINESTART":
            l.append(t)
    if ident_level != 0:  # If parsing end and ident_level is not 0, add as many DEDENT
        l.extend([Token("DEDENT", "")] * ident_level)
    if l[0].gettokentype() != "START":
        raise Exception("Wrong Start")
    return iter(l)  # Build an iterable with Token list to be used by parser


def lex(source):

    source = fix_idents(remove_empty_lines(remove_comments(source)))
    return off_side(lexer.lex(source))
