package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	arr := [5]int{2, 4, 6, 8, 10}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) { // анонимная функция с запуском горутины
			defer wg.Done()
			fmt.Println(arr[i] * arr[i])
		}(i) //определённое значение, передающееся анонимной функции
	}
	wg.Wait()
}

// // Альтернативная запись без передаваемого значения в анонимную функцию
// func main() {
// 	var wg sync.WaitGroup
// 	arr := [5]int{2, 4, 6, 8, 10}
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		k := i
// 		go func() { // анонимная функция с запуском горутины
// 			defer wg.Done()
// 			fmt.Println(arr[k] * arr[k])
// 		}() //определённое значение, передающееся анонимной функции
// 	}
// 	wg.Wait()
// }
