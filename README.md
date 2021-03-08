# Errorizer - A stringer, but for Errors

[![Go Reference](https://pkg.go.dev/badge/github.com/ryanfaerman/errorizer.svg)](https://pkg.go.dev/github.com/ryanfaerman/errorizer)

I really enjoy using Stringer on my constants, it sames me time and prevents
all sorts of errors in my code. Plus it's just neat!

This tool builds on that work (it uses most of the same code) but generates an
`Error` function instead. Others have written fantastic posts about the virtues
of constant errors, which my summarizing can't do justice. It also allows you
to have actual numerical error codes, which are useful for documentation and
long term support. Imagine being able to document Error-102 rather than just
the plain text.

That's the point of this tool.

I'm not even trying to pretend most of this code is my own, I've just modified
the Stringer tool to generate the Error function and change the string output
to the form: `package-CODE: error message`.

So, for a package named something like `sql` you'd get errors:

```
sql-1: invalid query
sql-2: missing something somewhere
sql-3: mismatched thing
```

See:

- https://dave.cheney.net/2016/04/07/constant-errors
- https://pkg.go.dev/golang.org/x/tools/cmd/stringer
