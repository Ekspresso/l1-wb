// Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)

func main() {
	sec := 3
	millisec := 3000
	microsec := 3000000
	min := 1
	h := 1
	mySleep(sec, "Second")
	mySleep(millisec, "Millisecond")
	mySleep(microsec, "Microsecond")
	mySleep(min, "Minute")
	mySleep(h, "Hour")
}

// Функция принимает первым параметром количество времени, а вторым единицы измерения времени в формате string.
// После чего останавливает выполнение функции, из которой вызывлась, на заданное время.
// Функция при отрицательном или нулевом значении первого аргумента сразу завершает свою работу.
// Второй аргумент:
// "Second" - секунды
// "Millisecond" - миллисекунды
// "Microsecond" - микросекунды
// "Minute" - минуты
// "Hour" - часы
//
func mySleep(d int, t string) {
	if d <= 0 {
		return
	}
	switch t {
	case "Second":
		<-time.After(time.Duration(d) * time.Second)
		fmt.Printf("%d seconds have passed\n", d)
	case "Millisecond":
		<-time.After(time.Duration(d) * time.Millisecond)
		fmt.Printf("%d milliseconds have passed\n", d)
	case "Microsecond":
		<-time.After(time.Duration(d) * time.Microsecond)
		fmt.Printf("%d microseconds have passed\n", d)
	case "Minute":
		<-time.After(time.Duration(d) * time.Minute)
		fmt.Printf("%d minutes have passed\n", d)
	case "Hour":
		<-time.After(time.Duration(d) * time.Hour)
		fmt.Printf("%d hours have passed\n", d)
	default:
		fmt.Println("Error: Incorrect time string")
	}
}
