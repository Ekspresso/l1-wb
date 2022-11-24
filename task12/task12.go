package main

import "fmt"

func main() {
	var strings = []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool)
	for i := 0; i < len(strings); i++ {
		set[strings[i]] = true
	}
	fmt.Println(set)
}
