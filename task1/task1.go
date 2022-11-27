// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

package main

import "fmt"

type Human struct {
	k, p int
	g    float64
	s    string
	//И ещё много других полей
}

//Возможные функции для структуры Human

func (h Human) RetStr(s string) string {
	return s + " from RetStr for Human"
}

func (h Human) Sum(k, p int) int {
	return k + p
}

//Реализация встраивания. Структура Action может использовать методы и поля структуры Human.
type Action struct {
	Human
	c int
	s string
}

func (a Action) RetStr(s string) string {
	return s + " ups from RetStr for action"
}

func (a Action) RetInt(c int) int {
	return c
}

// Примеры использования
func main() {
	a := Action{
		c: 3,
		s: "Hi, World from an Action!",
		Human: Human{
			k: 1,
			p: 2,
			g: 5.5,
			s: "Hello, world from a Human!"},
	}

	// При встраивании структур можно использовать методы встроенной структуры (Human) из Action,
	// а также напрямую обращаться к полям встроенной структуры
	fmt.Println("This is a sum: ", a.Sum(a.c, a.k))

	// Также можно без проблем использовать методы структуры Action, задействуя, в том числе, и поля встроенной структуры Human
	fmt.Println("This is p from Human: ", a.RetInt(a.p))

	// При наличии одноименных методов в Action и во встроенной Human структурах
	// при таком обращении будет иcпользоваться метод структуры Action.
	// То же касается и полей структур.
	fmt.Println(a.RetStr(a.s))

	// Если необходимо использовать одноимённое поле встроенной структуры, то к нему необходимо обращаться явно.
	fmt.Println(a.RetStr(a.Human.s))

	// Также явно необходимо обращаться и к одноимённому методу, если мы хотим использовать метод встроенной структуры.
	fmt.Println(a.Human.RetStr(a.Human.s))

	//Причём, если два встроенных класса будут иметь одноимённые поля или методы, то при неявном обращении к ним компилятор выдаст ошибку.
}
