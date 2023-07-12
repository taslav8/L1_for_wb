package main

import (
	"fmt"
	"strings"
)

func createHugeString(size int) string {
	return strings.Repeat("/", size)
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	buff := make([]byte, 100)
	copy(buff, v[:100])
	justString = string(buff)

	fmt.Println("Current string:", v)
	fmt.Println("The newest string:", justString)
}

func main() {
	someFunc()
}
