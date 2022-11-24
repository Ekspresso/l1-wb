//Программа принимает из 1 аргумента командной строки
//целое число, которое необходимо найти и выводит его номер
//в заданном массиве
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s [10]int
	for i := 0; i < len(s); i++ {
		s[i] = i + 1
	}
	if len(os.Args) > 2 {
		fmt.Println("Error: more than 1 search element")
	} else if len(os.Args) < 2 {
		fmt.Println("Error: the search element is not specified")
	} else {
		el, err := strconv.Atoi(os.Args[1])
		if err == nil {
			c := bin_search(el, s[:])
			if c == -1 {
				fmt.Println("Element not found or incorrect array")
			} else {
				fmt.Println(c)
			}
		}
	}
}

//Функция бинарного поиска получает элемент, который надо найти,
//срез отсортированного массива для поиска в нём и возвращает индекс найденного элемента.
func bin_search(el int, arr []int) int {
	l := 0
	r := len(arr)
	if arr[l] < arr[r-1] {
		for i := 0; i <= len(arr)/2; i++ {
			c := l + (r-l)/2
			if el < arr[c] {
				r = c
			} else if el > arr[c] {
				l = c
			} else {
				return (c)
			}
		}
	} else {
		for i := 0; i <= len(arr)/2; i++ {
			c := l + (r-l)/2
			if el < arr[c] {
				l = c
			} else if el > arr[c] {
				r = c
			} else {
				return (c)
			}
		}
	}
	return (-1)
}
