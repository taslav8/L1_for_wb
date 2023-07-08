package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func chanSol(array []int) {
	ch := make(chan int, 1)

	wg := sync.WaitGroup{}
	wg.Add(len(array))
	ch <- 0
	for _, value := range array {
		go func(val int, w *sync.WaitGroup, channel chan int) {
			channel <- (val * val) + <-channel
			wg.Done()
		}(value, &wg, ch)
	}
	wg.Wait()
	close(ch)
	fmt.Println(<-ch)
}

func mutexSol(array []int) {
	var sum int
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(array))
	for _, v := range array {
		go func(n int, m *sync.Mutex) {
			m.Lock()
			sum += n * n
			m.Unlock()
			wg.Done()
		}(v, &mu)
	}
	wg.Wait()
	fmt.Println(sum)
}

func atomicSol(array []int64) {
	var result int64
	wg := &sync.WaitGroup{}

	for _, v := range array {
		wg.Add(1)
		go func(num int64) {
			defer wg.Done()
			// потокобезопасная операция
			atomic.AddInt64(&result, num*num)
		}(v)
	}

	wg.Wait()

	fmt.Println(result)
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	mutexSol(arr)
	chanSol(arr)
	arr2 := make([]int64, len(arr))
	for i, v := range arr {
		arr2[i] = int64(v)
	}
	atomicSol(arr2)
}
