// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.

package main

import "fmt"

// В функции main объявляется переменная типа interface{}, а затем ей присваиваются различные значения разных типов.
// После чего эти значения передаются в функцию определения типов.
func main() {
	var p interface{}
	c := make(chan int)
	p = 32
	switchType(p)
	p = 21.5
	switchType(p)
	p = "Hello!"
	switchType(p)
	p = true
	switchType(p)
	p = nil
	switchType(p)
	p = c
	switchType(p)
	close(c)
	p = int64(31)
	switchType(p)
}

// Функция принимает переменную типа интерфейс и через switch case определяет тип переменной.
// После чего печатает информацию о том, переменная какого типа была получена.
// Для практического использования можно возвращать из этой функции различные значения
// и в вызывающей функции, в зависимости от вернувшегося результата, можно прописать определённые действия
// в зависимости от задачи.
func switchType(param interface{}) {
	switch p := param.(type) {
	case nil:
		fmt.Println("param is nil")
	case int:
		fmt.Println("param is int")
	case float64:
		fmt.Println("param is float64")
	case bool:
		fmt.Println("param is bool")
	case string:
		fmt.Println("param is string")
	case uint:
		fmt.Println("param is uint")
	case chan int:
		fmt.Println("param is channel")
	default:
		fmt.Printf("unexpected type of param: %T\n", p)
	}
}
