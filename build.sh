rm -rf parser/
java -jar ~/scripts/antlr4.jar -Dlanguage=Go -o parser Calc.g
go build Calc.go