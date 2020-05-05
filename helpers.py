import re


def remove_comments(source):
    comments = r"(#.*)(?:\n|\Z)"

    # Remove all comments
    comment = re.search(comments, source)
    while comment is not None:
        start, end = comment.span(1)
        source = source[0:start] + source[end:]  # remove string part that was a comment
        comment = re.search(comments, source)

    return source


def remove_empty_lines(source):
    multiline = r"([\s]+)(?:\n)"

    # Remove all empty lines
    line = re.search(multiline, source)
    while line is not None:
        start, end = line.span(1)
        source = (
            source[0:start] + source[end:]
        )  # remove string part that was an empty line
        line = re.search(multiline, source)

    return source


def fix_idents(source):
    idents = r"^[\s]+"
    notspaces = r"[^ ]"

    # Find ident length if spaces are used
    for oneline in source.split("\n"):
        ident = re.search(idents, oneline)  # Find first indented line
        if ident:
            ident_length = (
                ident.end()
            )  # First ident found (whitespace) define indent length.
            break

    # Replace spaces ident by tabs
    new_source = []
    for oneline in source.split("\n"):
        notspace = re.search(
            notspaces, oneline
        )  # Find first not space character
        if notspace:
            new_source.append(
                oneline[0 : notspace.start()].replace(" " * ident_length, "\t")
                + oneline[notspace.start() :]
            )  # Replace ident number of space with tabs
        else:
            new_source.append(oneline)

    return "\n".join(new_source)
