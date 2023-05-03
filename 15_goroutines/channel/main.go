// https://github.com/NerdCademyDev/golang/blob/main/21_goroutines/fibonacci/main.go
package main

import "fmt"

func fibonacci(max int, ch chan int) {
	fib := make([]int, max)
	fib[0] = 0
	fib[1] = 1
	ch <- fib[0]
	ch <- fib[1]

	for i := 2; i < max; i++ {
		fib[i] = fib[i-1] + fib[i-2]
		ch <- fib[i]
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go fibonacci(20, ch)

	for msg := range ch {
		fmt.Println(msg)
	}
}
