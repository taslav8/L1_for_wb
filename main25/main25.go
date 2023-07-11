package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	Sleep1(2 * time.Second)
	fmt.Println(time.Now().Sub(now))

	now = time.Now()
	Sleep2(2 * time.Second)
	fmt.Println(time.Now().Sub(now))

	now = time.Now()
	Sleep3(2 * time.Second)
	fmt.Println(time.Now().Sub(now))
}

func Sleep1(tm time.Duration) {
	<-time.After(tm)
}

func Sleep2(tm time.Duration) {
	tick := time.Tick(tm / 100)
	for i := 0; i < 100; i++ {
		<-tick
	}
}

func Sleep3(tm time.Duration) {
	timeout, cancelFunc := context.WithTimeout(context.Background(), tm)
	defer cancelFunc()
	<-timeout.Done()
}
