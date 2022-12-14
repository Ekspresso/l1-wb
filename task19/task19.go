// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

// Строка для теста: 1234😂⌘👍
// Передаётся аргументом при запуске программы
import (
	"fmt"
	"os"
	//"unicode/utf8"
)

// Способ 1. Программа принимает с запуском в качестве аргумента строку, которую в последствии переворачивает.
func main() {
	// Проверка на корректность переданных аргументов
	if len(os.Args) != 2 {
		fmt.Println("Error: Enter only 1 string")
		return
	}
	str := os.Args[1]
	rev := ""
	tempRev := ""
	// Цикл for range по срезу (string является срезом байт) вполняется по порядку.
	// Такой цикл, применительно к строкам, читает не байты по отдельности,
	// а руны (кодовые точки), которые могут состоять и из нескольких байт.
	// Обычный цикл с индексацией шёл бы только побайтово, что привело бы к потере целостности символов строки.
	for _, valRun := range str {
		rev = string(valRun)
		rev += tempRev
		tempRev = rev
	}
	fmt.Println(rev)
}

// Способ 2. Реализация с использованием библиотеки utf8.
// Функция utf8.DecodeRuneInString() принимает строку, затем побайтово считывает первую руну
// и возвращает 2 значения: прочитанную кодовую точку и количество считаных байт.
// Затем в цикле индекс i увеличивается на количество прочитанных байт,
// чтобы затем начать считывать с этого места.
// Далее наращивания строк происходит как в способе 1.
// func main() {
// 	if len(os.Args) != 2 {
// 		fmt.Println("Error: Enter only 1 string")
// 		return
// 	}
// 	str := os.Args[1]
// 	rev, tempRev := "", ""
// 	for i, w := 0, 0; i < len(str); i += w {
// 		valRun, width := utf8.DecodeRuneInString(str[i:])
// 		w = width
// 		rev = string(valRun)
// 		rev += tempRev
// 		tempRev = rev
// 	}
// 	fmt.Println(rev)
// }
