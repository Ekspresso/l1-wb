// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….)
// с использованием конкурентных вычислений.

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
	k = k * k
	mutex.Lock() // Блокировка мьютекса для работы с общей переменной для записи "sum"
	*sum += k
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
// 			k := arr[i] * arr[i]
// 			mutex.Lock()
// 			sum += k
// 			mutex.Unlock()
// 		}(i)
// 	}
// 	wg.Wait()
// 	fmt.Println(sum)
// }
