// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

// Строка для теста: "1234😂⌘👍 cool👍 lol😂 hello"
// Передаётся именно в кавычках аргументом при запуске программы,
// чтобы эта строка воспринималась как один аргумент.

import (
	"fmt"
	"os"
)

// Способ 1. Программа принимает с запуском в качестве аргумента строку, в которой в последствии переворачивает слова.
// Порядок следования символов в слове при этом остаётся неизменным.
func main() {
	// Проверка на корректность переданных аргументов
	if len(os.Args) != 2 {
		fmt.Println("Error: Enter only 1 string")
		return
	}
	str := os.Args[1]
	rev := ""
	tempRev := ""
	tempWord := ""
	// Цикл for range по срезу (string является срезом байт) вполняется по порядку.
	// Такой цикл, применительно к строкам, читает не байты по отдельности,
	// а руны (кодовые точки), которые могут состоять и из нескольких байт.
	// Обычный цикл с индексацией шёл бы только побайтово, что привело бы к потере целостности символов строки.
	for _, valRun := range str {
		// Если встреченный символ не пробел, то слово просто считывается в переменную tempWord
		if string(valRun) != " " {
			tempWord += string(valRun)
		} else {
			// Если встреченный символ является пробелом, то выполняется запись слов с переворотом.
			rev = tempWord + string(valRun)
			rev += tempRev
			tempRev = rev
			tempWord = ""
		}
	}
	// Если в переменной временного хранения слов tempWord ещё есть данные, то они дописываются в начало строки.
	if tempWord != "" {
		rev = tempWord + " "
		rev += tempRev
	}
	fmt.Println(rev)
}

// Способ 2. Реализация переворота слов в строке с использованием функции Split
// из библиотеки strings, где в качестве разделителя используется пробел.
// func main() {
// 	// Проверка на корректность переданных аргументов
// 	if len(os.Args) != 2 {
// 		fmt.Println("Error: Enter only 1 string")
// 		return
// 	}
// 	str := os.Args[1]
// 	rev := ""
// 	tempRev := ""
// 	words := strings.Split(str, " ")
// 	for _, word := range words {
// 		if word != "" {
// 			rev = word + " "
// 			rev += tempRev
// 			tempRev = rev
// 		}
// 	}
// 	fmt.Println(rev)
// }
