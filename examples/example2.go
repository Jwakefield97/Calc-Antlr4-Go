package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"../parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type calcListener struct {
	*parser.BaseCalcListener

	stack []int
}

func (l *calcListener) push(i int) {
	l.stack = append(l.stack, i)
}

func (l *calcListener) pop() int {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func (l *calcListener) ExitMulDiv(c *parser.MulDivContext) {
	right, left := l.pop(), l.pop()

	switch c.GetOp().GetTokenType() {
	case parser.CalcParserMUL:
		l.push(left * right)
	case parser.CalcParserDIV:
		l.push(left / right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *calcListener) ExitAddSub(c *parser.AddSubContext) {
	right, left := l.pop(), l.pop()

	switch c.GetOp().GetTokenType() {
	case parser.CalcParserADD:
		l.push(left + right)
	case parser.CalcParserSUB:
		l.push(left - right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *calcListener) ExitNumber(c *parser.NumberContext) {
	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		panic(err.Error())
	}

	l.push(i)
}

// calc takes a string expression and returns the evaluated result.
func calc(input string) int {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalcParser(stream)

	// Finally parse the expression (by walking the tree)
	var listener calcListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	return listener.pop()
}

func main() {
	// // Setup the input
	// is := antlr.NewInputStream("1 + 2 * 3")

	// // Create the Lexer
	// lexer := parser.NewCalcLexer(is)
	// stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// // Create the Parser
	// p := parser.NewCalcParser(stream)

	// // Finally parse the expression
	// antlr.ParseTreeWalkerDefault.Walk(&calcListener{}, p.Start())
	fmt.Println("Common calculator operations are supported (Add/Sub/Mult/Div). Type exit to leave the application.")
	for {
		buf := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		line, err := buf.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		} else {
			if strings.EqualFold(line, "exit\n") {
				fmt.Println("exiting...")
				break
			} else {
				fmt.Println("= " + strconv.Itoa(calc(line)))
			}
		}
	}
}
