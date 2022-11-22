package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var sum int
	var mutex sync.Mutex // Мьютекс нужен для безопасного доступа к памяти несколькими горутинами

	arr := [5]int{2, 4, 6, 8, 10}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go sumSq(&sum, arr[i], &mutex, &wg)
	}
	wg.Wait()
	fmt.Println(sum)
}

// Функция принимает по указателю "sum" и прибавляет к её значению квадрат переданного чила k
func sumSq(sum *int, k int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock() // Блокировка мьютекса для работы с общей переменной для записи "sum"
	*sum += k * k
	mutex.Unlock() // Разблокировка мьютекса
}

// // Альтернативная версия написания с использованием анонимной функции
// func main() {
// 	var wg sync.WaitGroup
// 	var sum int
// 	var mutex sync.Mutex
//
// 	arr := [5]int{2, 4, 6, 8, 10}
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			mutex.Lock()
// 			sum += arr[i] * arr[i]
// 			mutex.Unlock()
// 		}(i)
// 	}
// 	wg.Wait()
// 	fmt.Println(sum)
// }
