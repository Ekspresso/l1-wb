// Удалить i-ый элемент из слайса.

package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7}
	sl = deleteElSlise(sl, 2)
	fmt.Println(sl)
}

// Функция добавляет к срезу от начала до k-ого элемента (не включительно) срез от следующего после k элемента до конца.
func deleteElSlise(sl []int, k int) []int {
	sl[k] = 0
	return append(sl[:k], sl[(k+1):]...)
}
