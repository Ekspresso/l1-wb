// Разработать программу, которая перемножает, делит, складывает,
// вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
	"os"
)

// Программа принимает в качестве аргументов с запуском 3 строки: первое число, знак операции и второе число.
func main() {
	// Проверка на корректность переданных аргументов
	if len(os.Args) != 4 {
		fmt.Println("Error: Enter: arg1:number arg2:sign arg3:number")
		return
	}
	if check(os.Args[1], os.Args[2], os.Args[3]) {
		return
	}

	// Переменные хранятся в big int для корректной работы с большими числами.
	sign := os.Args[2]
	res := new(big.Int)
	a := new(big.Int)
	a.SetString(os.Args[1], 10)
	b := new(big.Int)
	b.SetString(os.Args[3], 10)
	// Выполнение операции в зависимости от введённого знака.
	switch sign {
	case "*":
		res.Mul(a, b)
		fmt.Println(res)
	case "+":
		res.Add(a, b)
		fmt.Println(res)
	case "/":
		if b.Sign() == 0 {
			fmt.Println("Error: b is 0")
			return
		}
		res.Div(a, b)
		fmt.Println(res)
	case "-":
		res.Sub(a, b)
		fmt.Println(res)
	}
}

// Функция проверки на корректность введённых данных.
func check(a, sign, b string) bool {
	for _, symb := range a {
		if symb < '0' || symb > '9' {
			fmt.Println("Eror: 1 arg not a number")
			return true
		}
	}
	for _, symb := range b {
		if symb < '0' || symb > '9' {
			fmt.Println("Eror: 3 arg not a number")
			return true
		}
	}
	for ind, symb := range sign {
		if ind > 0 {
			fmt.Println("Eror: more than 1 symb in sign")
			return true
		}
		if symb != '+' && symb != '-' && symb != '/' && symb != '*' {
			fmt.Println("Error: 2 arg not a sign")
			return true
		}
	}
	return false
}
