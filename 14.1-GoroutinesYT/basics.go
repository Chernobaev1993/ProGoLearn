package main

import (
	"fmt"
	"sync"
)

func demoGorouts() {
	var wg sync.WaitGroup	
	total := 0
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(index int) {
			ind, tot := index, total
			total += index
			fmt.Println(ind, tot, total)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("\nTotal =", total)
}
