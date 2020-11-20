package main

import (
	"fmt"
	"sync"
)

func main() {
	bargains := make(chan int, 10)

	for i := 0; i < 10; i++ {
		bargains <- i
	}

	var wg sync.WaitGroup
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(num int) {
			select {
			case result := <-bargains:
				fmt.Printf("User: %d get %d\n", num, result)
			default:
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
