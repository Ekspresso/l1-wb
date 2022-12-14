// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

package main

import (
	"fmt"
	// "unicode"
)

func main() {
	s := ";e mp[😂⌘"
	fmt.Println(checkUniqStr(s))
}

// Функция проверки строку на уникальность символов в ней.
// Принимает строку, возвращает false, если  в строке есть повторяющиеся символы и true если нет.
// Функция регистронезависимая, то есть если будут содержаться две одинаковые буквы одного языка (английского или русского),
// но с разными регистрами, то функция вернёт false.
func checkUniqStr(str string) bool {
	m := make(map[rune]bool)
	for _, symb := range str {
		if m[symb] {
			return false
		}
		if symb >= 'a' && symb <= 'z' && m[symb-32] {
			return false
		}
		if symb >= 'A' && symb <= 'Z' && m[symb+32] {
			return false
		}
		if symb >= 'А' && symb <= 'Я' && m[symb+32] {
			return false
		}
		if symb >= 'а' && symb <= 'я' && m[symb-32] {
			return false
		}
		m[symb] = true
	}
	return true
}

// Реализация функци проверки с использованием пакета unicode.
// func checkUniqStr(str string) bool {
// 	m := make(map[rune]bool)
// 	for _, symb := range str {
// 		symb = unicode.ToLower(symb)
// 		if m[symb] {
// 			return false
// 		}
// 		m[symb] = true
// 	}
// 	return true
// }
