rm -rf parser/
antlr -Dlanguage=Go -o parser Calc.g # assumes you have antlr as an alias to the antlr4 jar.
go build Calc.go