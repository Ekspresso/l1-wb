package main

import "fmt"

func main() {
	// Создание карт для хранения множеств
	m1 := make(map[int]bool)
	m2 := make(map[int]bool)
	interect := make(map[int]bool)
	// Заполнение множества m1
	for i := 3; i < 10; i++ {
		m1[i] = true
	}
	// Заполнение множества m2
	for i := 0; i < 15; i += 2 {
		m2[i] = true
	}
	// Запись пересечения множеств m1 и m2 в множество interect
	for key := range m1 {
		if m2[key] {
			interect[key] = true
		}
	}
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(interect)
}
