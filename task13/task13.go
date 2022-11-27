// Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	a := 1
	b := 100
	fmt.Println(a, " ", b)
	// Меняем местами переменные a и b
	a, b = b, a
	fmt.Println(a, " ", b)
}
