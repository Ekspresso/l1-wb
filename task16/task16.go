package main

import "fmt"

func main() {
	arr := [5]int{1, 0, 5, 11, 9}
	myQuickSort(arr[:], 0, len(arr)-1)
	fmt.Println(arr)
}

// Рекурсивная функция для сортировки.
// сначала выбирается pivot - опорный элемент, относительно которого сортируется массив.
// В этой реализации опорный элемент всегда справа.
// Затем все элементы, которые меньше опорного записываются с самого начала массива.
// После чего сам pivot помещается в место, где заканчиваются все элементы меньше него.
// После чего массив разбивается пополам и все действия повторяются для каждой половины.
// Так происходит до момента, пока не будут рассматриваться массивы из 2 элементов.
// В конце всех манипуляций массив, который был передан в программу становится отсортированным,
// так как он передавался в функцию через срез и реалокации памяти не происходило.
func myQuickSort(arr []int, left int, right int) {
	if left < right {
		pivot := partition(arr, left, right)
		myQuickSort(arr, left, pivot-1)
		myQuickSort(arr, pivot+1, right)
	}
}

func partition(arr []int, left int, right int) int {
	pivot := arr[right]
	i := left - 1
	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			i++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[right], arr[i+1] = arr[i+1], arr[right]
	return (i + 1)
}
