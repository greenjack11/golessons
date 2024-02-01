package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const NAN = -11

// this is a comment
func add(x int, y int) int {
	return x + y
}

func substract(x int, y int) int {
	return x - y
}

func multiply(x int, y int) int {
	return x * y
}

func devide(x int, y int) int {
	return x / y
}

func romanToInt(x string) int {
	switch x {
	case "I":
		return 1
	case "II":
		return 2
	case "III":
		return 3
	case "IV":
		return 4
	case "V":
		return 5
	case "VI":
		return 6
	case "VII":
		return 7
	case "VIII":
		return 8
	case "IX":
		return 9
	case "X":
		return 10
	}
	return NAN
}

func intToRoman(n int) string {
	i := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	x := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	var str string
	if n > 100 || n < 1 {

		panic("The number can only be from 1 to 100")
	}
	if n == 100 {
		return "C"
	}
	x_idx := n / 10
	n = n % 10
	str = str + x[x_idx]
	str = str + i[n]

	return str
}

func isArabic(x string) bool {
	for i := 0; i < len(x); i++ {
		if x[i] < '0' || x[i] > '9' {
			return false
		}
	}
	return true
}

func isRoman(x string) bool {
	for i := 0; i < len(x); i++ {
		if !(x[i] == 'I' || x[i] == 'X' || x[i] == 'V') {
			return false
		}
	}
	return true
}

type Expression struct {
	a, b     int
	operator byte
	isRoman  bool
}

func (expr Expression) evaluate() string {
	var res int
	switch expr.operator {
	case '*':
		res = multiply(expr.a, expr.b)
	case '/':
		res = devide(expr.a, expr.b)
	case '+':
		res = add(expr.a, expr.b)
	case '-':
		res = substract(expr.a, expr.b)
	}
	if expr.isRoman {
		return intToRoman(res)
	}
	return strconv.Itoa(res)
}

func tokenize(str string) Expression {
	signs := [4]string{"+", "-", "/", "*"}
	var terms [2]string
	for _, sign := range signs {
		tokens := strings.Split(str, sign)
		if len(tokens) == 2 {
			for i, str := range tokens {
				terms[i] = strings.Trim(str, "+-/* ")
			}
			var a, b int
			if isRoman(terms[0]) && isRoman(terms[1]) {

				a = romanToInt(terms[0])
				b = romanToInt(terms[1])
				if a != NAN && b != NAN {
					res := Expression{a, b, byte(sign[0]), true}
					return res
				} else {
					panic("Error: one of numbers is invalid")
				}
			} else {
				if isArabic(terms[0]) && isArabic(terms[1]) {

					a, e1 := strconv.ParseInt(terms[0], 10, 32)
					b, e2 := strconv.ParseInt(terms[1], 10, 32)

					if e1 == nil && e2 == nil {
						res := Expression{int(a), int(b), byte(sign[0]), true}
						return res
					} else {
						panic("Error: one of numbers is invalid")
					}

				} else {
					panic("Error: Arabic and roman numbers have been mixed, or input string contain not a number.")
				}
			}
		}
	}
	res := Expression{NAN, NAN, 0, false}
	return res
}

func main() {
	var text string
	text = ""
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text = scanner.Text()
	}
	//var terms [2]int
	expr := tokenize(text)
	res := expr.evaluate()
	fmt.Println(res)
}
