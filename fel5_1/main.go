/*
1. Tervezzen algoritmust és írjon olyan programot, ami fordított lengyel formájú
kifejezéseket értékel ki, verem segítségével! Fordított lengyel formán egy
kifejezésnek olyan formáját értjük, amelyben számok illetve műveleti jelek
szerepelnek egymástól szóközzel elválasztva. A kiértékelés úgy történik, hogy
amennyiben a sorozatban szám következik azt betesszük a verembe, amennyiben
művelet, a verem tetején lévő két értékkel elvégezzük a műveletet, és az eredmény
a verembe kerül. Például a 3*(4-5*(3-15)/4)+100 kifejezés fordított lengyel formájú
alakja 100 15 3 - 4 5 / * 4 - 3 * + Tesztelje a programot a 2 5 + 3 * 9 2 1 + / -
kifejezéssel!

1. Design an algorithm and write a program that evaluates expressions in Reverse Polish Notation using a stack! By Reversed Polish Notation we mean a form of an expression in which numbers and operation signs are included, separated from each other by spaces. The evaluation is done in such a way that if there is a number in the sequence, it is placed on the stack, if it is an operation, the operation is performed with the two values at the top of the stack, and the result is placed on the stack. For example, the expression 3*(4-5*(3-15)/4)+100 has the reverse Polish Notation of
100 15 3 - 4 5 / * 4 - 3 * +
Test the program with this expression: 2 5 + 3 * 9 2 1 + / -

evaluation of test expression:
2 5 + 3 * 9 2 1 + / -
7 3 * 9 2 1 + / -
21 9 2 1 + / -
21 9 3 / -
21 3 -
18
*/
package main

import (
	"fmt"
	"log"
	"strconv"
)

type tokenType = string

const (
	num tokenType = "num"
	add tokenType = "+"
	sub tokenType = "-"
	div tokenType = "/"
	mul tokenType = "*"
)

type token struct {
	type_  tokenType
	numVal float64
}

func main() {
	input := "2 5 + 3 * 9 2 1 + / -"
	fmt.Printf("Input: %s\n\n", input)
	tokens := scan(input)
	res := eval(tokens)
    fmt.Printf("Output: %g", res)
}

func eval(tokens []token) float64 {
	var stack []float64
	i := 0
	for _, t := range tokens {
		switch t.type_ {
		case num:
			stack = append(stack, t.numVal)
			i++
		case add:
			stack[i-2] = stack[i-2] + stack[i-1]
            stack = stack[:i-1]
			i--
		case sub:
			stack[i-2] = stack[i-2] - stack[i-1]
            stack = stack[:i-1]
			i--
		case div:
			stack[i-2] = stack[i-2] / stack[i-1]
            stack = stack[:i-1]
			i--
		case mul:
			stack[i-2] = stack[i-2] * stack[i-1]
            stack = stack[:i-1]
			i--
		}
	}
	return stack[0]
}

func scan(text string) []token {
	var tokens []token
	startOfNum := -1
	for i, c := range text {
		if '0' <= c && c <= '9' || c == '.' {
			if startOfNum == -1 {
				startOfNum = i
			}
		} else if startOfNum != -1 {
			numLexeme := text[startOfNum:i]
			numVal, err := strconv.ParseFloat(numLexeme, 64)
			if err != nil {
				log.Fatalf("there was a badly formated number in the input at pos %d. Number tried to be converted: %s", i, numLexeme)
			}
			tokens = append(tokens, token{type_: num, numVal: numVal})
			startOfNum = -1
		}
		switch c {
		case '/':
			tokens = append(tokens, token{type_: div})
		case '*':
			tokens = append(tokens, token{type_: mul})
		case '-':
			tokens = append(tokens, token{type_: sub})
		case '+':
			tokens = append(tokens, token{type_: add})
		}
	}
	return tokens
}
