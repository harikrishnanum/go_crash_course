package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Printf("Goroutine %d executing\n", id)
			time.Sleep(time.Second)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("All Goroutines completed")
}
