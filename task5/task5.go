package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Создание канала, счётчика процессов wg и контекста с отменой через заданное время.
	// Контекст позволяет корректно завершить работу программы.
	c := make(chan int)
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Вызов горутины для отслеживания системных сигналов на прекращение работы программы.
	wg.Add(1)
	go chanSys(c, ctx, cancel, &wg)
	// Вызов горутины-воркера которая читает данные из канала и выводит их в стандартный вывод.
	wg.Add(1)
	go worker(c, &wg, ctx)

	// Бесконечный цикл, который последовательно записывает данные в канал.
	// Данные представляют из себя цифры от 0 до 9 включительно.
	for {
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				// Если контекст отменён, то закрывает канал, ожидает завершения всех горутин и завершает работу программы.
				fmt.Println("Closing channel")
				close(c)
				wg.Wait()
				fmt.Println("All gorutins have ended. Exit.")
				return
			default:
				// В штатном режиме выводит информацию о том, что происходит запись в канал
				// и передаёт в него данные (цифру).
				fmt.Printf("Writing data to a channel. Data: %d\n", i)
				time.Sleep(time.Second)
				c <- i
			}
		}
	}
}

// Функция воркер. Читает данные из канала и печатает их в стандартный вывод с текстовыми разъяснениями.
// Работает в бесконечном цикле до отмены контекста.
func worker(c chan int, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done() // Уменьшение счётчика процессов на 1 при завершении работы функции.
	for {
		select {
		case <-ctx.Done():
			// Если контекст отменён, то печатает информацию о завершении работы воркера и завершает работу функции.
			fmt.Printf("Ending Worker\n")
			return
		default:
			// В штатном состоянии принимает данные из канала с проверкой канала на открытость.
			// Если канал открыт, то печатает полученные данные с текстовыми разъяснениями.
			// Если канал закрыт, то печатает информацию об этом и завершает работу функции.
			time.Sleep(time.Second)
			data, ok := <-c
			if !ok {
				fmt.Printf("Worker: channel is closed. Ending Worker\n")
				return
			}
			fmt.Printf("Data from worker: %d\n", data)
		}
	}
}

// Функция отслеживающая системный канал, в который поступают сигналы из вне.
// Также отслеживает канал контекста.
// Функция реагирует на сигнал завершения работы программы и отменяет контекст,
// что передаётся в другие функции, которые следят за контекстом.
// Если контекст был отменён в другом месте, то завершает работу функции.
func chanSys(c chan int, ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшение счётчика процессов на 1 при завершении работы функции.
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		// Завершение работы функции при отменённом контексте.
		return
	case <-signalCh:
		// Отмена контекста при получении сигнала и завершение работы функции.
		cancel()
		return
	}
}
