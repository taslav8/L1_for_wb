package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var a, b int64 = 999, 121
	fmt.Printf("До: a = %d, b = %d\n", a, b)

	//******************************************** Способы:
	a, b = b, a
	fmt.Printf("После:  a = %d, b = %d\n", a, b)

	a ^= b
	b ^= a
	a ^= b
	fmt.Printf("После:  a = %d, b = %d\n", a, b)

	b = atomic.SwapInt64(&a, b)
	fmt.Printf("После:  a = %d, b = %d\n", a, b)

}
