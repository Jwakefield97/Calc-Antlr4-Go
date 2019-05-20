package main

import (
	//"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"io/ioutil"
	"./parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type calcListener struct {
	*parser.BaseCalcListener
	Variables map[string]float64
	stack []float64
}

func (l *calcListener) push(i float64) {
	l.stack = append(l.stack, i)
}

func (l *calcListener) pop() float64 {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

//multiplication or division used in an expression
func (l *calcListener) ExitMulDiv(c *parser.MulDivContext) {
	right, left := l.pop(), l.pop()

	switch c.GetOp().GetTokenType() {
	case parser.CalcParserMUL:
		l.push(left * right)
	case parser.CalcParserDIV:
		l.push(left / right)
	case parser.CalcParserCARAT:
		l.push(math.Pow(left,right))
	case parser.CalcParserMODULO: 
		l.push(float64(int(left) % int(right)))
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

//addition or subtraction in an expression
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

//number used in an expression
func (l *calcListener) ExitNumber(c *parser.NumberContext) {
	i, err := strconv.ParseFloat(c.GetText(), 64)
	if err != nil {
		panic(err.Error())
	}

	l.push(i)
}

//variable used in an expression
func (l *calcListener) ExitVariableExp(c *parser.VariableExpContext) {
	l.push(l.Variables[c.GetText()])
}

//printing an expression
func (l *calcListener) ExitPrintExp(c *parser.PrintExpContext) {
	exp := fmt.Sprintf("%f",l.pop())
	fmt.Println(exp)
}

//printing a variable
func (l *calcListener) ExitPrintVar(c *parser.PrintVarContext) {
	exp := fmt.Sprintf("%f",l.pop())
	fmt.Println(exp)
}

//get the name of the variable from the declaration
func (l *calcListener) getVarName(dec string) string {
	noLet := strings.Replace(dec,"let","",-1);
	return strings.TrimSpace(strings.Split(noLet,"=")[0])
}

//variable declaration
func (l *calcListener) ExitVariable(c *parser.VariableContext) {
	varName := l.getVarName(c.GetText());
	i, _ := strconv.ParseFloat(c.GetStop().GetText(), 64)
	l.Variables[varName] = i
}

func NewCalcListener() calcListener {
	c := calcListener{}
	c.Variables = map[string]float64{}
	return c
}

// calc takes a string expression and returns the evaluated result.
func calc(input string) {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalcParser(stream)

	// Finally parse the expression (by walking the tree)
	listener := NewCalcListener()
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())
}

func main() {
	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
        fmt.Print(err)
		panic("error opening file")
    }
	calc(string(bytes))
	// fmt.Println("Common calculator operations are supported (Add/Sub/Mult/Div). Type exit to leave the application.")
	// for {
	// 	buf := bufio.NewReader(os.Stdin)
	// 	fmt.Print("> ")
	// 	line, err := buf.ReadString('\n')
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		if strings.EqualFold(line, "exit\n") {
	// 			fmt.Println("exiting...")
	// 			break
	// 		} else {
	// 			calc(line)
	// 		}
	// 	}
	// }
}
