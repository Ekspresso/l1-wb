package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex // Мьютекс нужен для безопасного доступа к памяти несколькими горутинами.
	m := make(map[int]int)

	arr := [5]int{2, 4, 6, 8, 10}
	// Запуск горутин для конкурентной записи в map данных из массива.
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go mapWrite(i, m, arr[i], &mutex, &wg)
	}
	// Ожидание завершения запущенных горутин.
	wg.Wait()

	for key, val := range m {
		fmt.Printf("%d : %d\n", key, val)
	}
}

// Функция принимает по указателю "map" и записывает в неё новые значения
func mapWrite(i int, m map[int]int, k int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock() // Блокировка мьютекса для записи в map
	fmt.Printf("Writing: %d:%d\n", i, k)
	m[i] = k
	mutex.Unlock() // Разблокировка мьютекса
}
