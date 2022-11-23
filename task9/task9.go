package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	x := make(chan int)
	x2 := make(chan int64)

	arr := [5]int{2, 4, 6, 8, 10}

	// Горутина, записывающая данные в канал 1
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < len(arr); i++ {
			// fmt.Printf("Sending %d to ch1\n", arr[i])
			x <- arr[i]
		}
		close(x)
	}()

	// Горутина, записывающая число, полученное из канала 1, в канал 2 с возведением числа в квадрат.
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			datax, ok := <-x
			if !ok {
				close(x2)
				return
			}
			// fmt.Printf("Reading %d from ch1 and sending %d^2 to ch2\n", datax, datax)
			x2 <- int64(datax) * int64(datax)
		}
	}()

	// Горутина, читающая данные из канала 2 и выводящая их в stdout.
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			datax2, ok := <-x2
			if !ok {
				return
			}
			fmt.Printf("Data from ch2: %d\n", datax2)
		}
	}()

	// Ожидание завершения запущенных горутин.
	wg.Wait()
	fmt.Println("End of programm")
}
