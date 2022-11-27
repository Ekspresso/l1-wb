// Реализовать все возможные способы остановки выполнения горутины.

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

// Функция main содержит набор горутин, который представляет различные способы завершения работы горутин.
// В основном все приведённые методы завершения работы основываются на чтении из какого-либо канала.
func main() {
	// Канал для передачи сообщения
	c := make(chan bool)
	// Канал для закрытия
	a := make(chan int)
	// Счётчик количества процессов
	var wg sync.WaitGroup
	// Контекст для отмены по истечению времени
	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second*2)
	// Контекст для отмены программно в явном виде
	ctx2, cancel2 := context.WithCancel(context.Background())
	// Контекст для отмены при сигнале завершения работы программы
	ctxAll, cancelAll := context.WithCancel(context.Background())
	defer cancel1()
	defer cancel2()
	defer cancelAll()

	// Первый метод прекращения горутин - это сигнал на завершение работы программы.
	// При поступлении такого сигнала происходит отмена контекста и все горутины,
	// которые следят за ним, получают отмену контекста для дальнейшего завершения.
	wg.Add(1)
	go chanSys(ctxAll, cancelAll, &wg)

	// Второй метод прекращения работы горутин - это отмена контекста по истечению времени.
	// Горутина, помимо основных действий, прослушивает канал контекста, и если контекст отменён по истечению времени,
	// то происходит завершение работы горутины.
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx1.Done():
				fmt.Println("Ending 1 by context time out")
				return
			case <-ctxAll.Done():
				fmt.Println("Ending 1 by context sys signal")
				return
			default:
			}
		}
	}()

	// Третий способ прекращения работы горутин - это отмена контекста в коде самой программы.
	// Горутина, помимо основных действий, прослушивает канал контекста, и если контекст был отменён,
	// то происходит завершение работы горутины.
	// Отмена контекста может происходить в разных ситуациях, и, при необходимости, есть возможность следить за ним
	// и завершать работу горутин.
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx2.Done():
				fmt.Println("Ending 2 by context user cancel")
				return
			case <-ctxAll.Done():
				fmt.Println("Ending 2 by context sys signal")
				return
			default:
			}
		}
	}()

	// Четвёртый способ прекращения работы горутин - это поступление сообщения из определённого канала.
	// В данном случае горутина выполняет все описанные в коде действия для этого канала и завершает свою работу.
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-c:
				fmt.Println("Ending 3 by chanel message")
				close(c)
				return
			case <-ctxAll.Done():
				fmt.Println("Ending 3 by context sys signal")
				return
			default:
			}
		}
	}()

	// Пятый способ - это прекращение работы горутины при закрытии канала.
	// При чтении из канала горутина проверяет открыт ли он, и если канал закрыт, то завершает свою работу.
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctxAll.Done():
				fmt.Println("Ending 4 by context sys signal")
				return
			default:
				data, ok := <-a
				if !ok {
					fmt.Println("Ending 4 by chanel closing")
					return
				}
				_ = data
			}
		}
	}()

	// Шестой способ - это завершение работы горутины естественным образом.
	// Горутина просто выполняет все прописанные действия и завершает свою работу.
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Ending 5 on completion of work")
	}()

	// Отмена контекста в коде (явно), приводящая к завершению работы горутины.
	cancel2()

	// Передача сообщения в канал, приводящая к завершению работы горутины.
	c <- true

	// Закрытие канала, приводящее к завершению работы горутины.
	close(a)

	fmt.Println("To exit from program: ctrl+c")
	// Ожидание завершения работы всех горутин.
	wg.Wait()
	fmt.Println("All gorutins have ended. Exit.")
}

// Функция отслеживающая системный канал, в который поступают сигналы из вне.
// Функция реагирует на сигнал завершения работы программы и отменяет контекст,
// что передаётся в другие функции, которые следят за контекстом.
func chanSys(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup) {
	defer wg.Done()
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	<-signalCh
	fmt.Println("Ending all with ctrl+c")
	cancel()
}
