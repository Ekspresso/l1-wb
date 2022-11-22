package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Обработка переданного аргумента с запуском программы. Аргументом передаётся количество воркеров.
	if len(os.Args) > 2 || len(os.Args) < 2 {
		fmt.Println("Error: enter only the number of workers using the argument")
		return
	}

	// Конвертация количества воркеров из строки в число (Args[i] тип string).
	k, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// Создание канала, счётчика процессов wg и контекста с функцией отмены.
	// Контекст позволяет корректно завершить работу программы.
	c := make(chan int)
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Вызов горутины для отслеживания системных сигналов на прекращение работы программы.
	go chanSys(ctx, cancel)

	// Создание заданного количества воркеров с увеличением счётчика процессов на один при каждом создании.
	for i := 0; i < k; i++ {
		wg.Add(1)
		go worker(c, i, &wg, ctx)
	}

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
func worker(c chan int, i int, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done() // Уменьшение счётчика процессов на 1 при завершении работы функции.
	for {
		select {
		case <-ctx.Done():
			// Если контекст отменён, то печатает информацию о завершении работы воркера и завершает работу функции.
			fmt.Printf("Ending Worker %d\n", i)
			return
		default:
			// В штатном состоянии принимает данные из канала с проверкой канала на открытость.
			// Если канал открыт, то печатает полученные данные с текстовыми разъяснениями.
			// Если канал закрыт, то печатает информацию об этом и завершает работу функции.
			time.Sleep(time.Second)
			data, ok := <-c
			if !ok {
				fmt.Printf("Worker %d: channel is closed. Ending Worker %d\n", i, i)
				return
			}
			fmt.Printf("Data from worker %d: %d\n", i, data)
		}
	}
}

// Функция отслеживающая системный канал, в который поступают сигналы из вне.
// Функция реагирует на сигнал завершения работы программы и отменяет контекст,
// что передаётся в другие функции, которые следят за контекстом.
func chanSys(ctx context.Context, cancel context.CancelFunc) {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	<-signalCh
	cancel()
}
