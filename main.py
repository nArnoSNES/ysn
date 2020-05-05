import sys
import lexer


def main():
    source = sys.stdin.read()
    try:
        toks = lexer.lex(source)
    except Exception as e:
        sys.exit("Oops! %s" % str(e))
    for t in toks:
        print(t)


if __name__ == "__main__":
    main()
