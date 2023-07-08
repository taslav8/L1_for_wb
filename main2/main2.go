package main

import (
	"fmt"
	"sync"
)

func calculate(arr []int) {
	var wg sync.WaitGroup

	for _, n := range arr {
		wg.Add(1)

		go func(i int) {
			fmt.Println(i * i) //Расчет квадрата и вывод его в отдельной горутине
			wg.Done()
		}(n)
	}

	wg.Wait()

}

func main() {
	arr := []int{2, 4, 6, 8, 10}

	calculate(arr)
}

//Вывод ответов будет в случайном порядке
