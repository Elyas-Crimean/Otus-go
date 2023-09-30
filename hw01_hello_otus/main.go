package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

// Инвертирование строки с использованием стороннего модуля
func main() {
	const source = "Hello, OTUS!"
	result := stringutil.Reverse(source)
	fmt.Println(result)
}
