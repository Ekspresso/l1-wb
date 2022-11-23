package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Enter two numbers: the number and the bit to change")
		return
	}
	numb, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	i, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	var res int64
	ibit := int64(1) << i // Получение числа int64, битовое представление которого будет на подобии "10000...", где 1 указывает на i-ый бит.
	if numb&ibit == 0 {   // Если i-ый бит в numb равен 0, то результат операции будет 0 (логическое "И").
		res = numb | ibit // Логическое "ИЛИ". Бит равен 1, если хотя бы один из сравниваемых битов равен 1. Бит равен 0, ели оба равны 0.
	} else {
		res = numb &^ ibit //"И НЕ" Бит res равен 0, если соответствующий бит ibit равен 1. Если бит в ibit равен 0, то берется значение соответствующего бита из numb.
	}
	fmt.Println(numb)
	fmt.Println(strconv.FormatInt(numb, 2))
	fmt.Println(strconv.FormatInt(res, 2))
	fmt.Println(res)
}
