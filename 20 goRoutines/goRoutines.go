package main

import "sync"

// "time"

func task1(a int, w *sync.WaitGroup) {
	defer w.Done()
	println("task1", a)
}



func main() {
	var wg sync.WaitGroup
	for i := 0; i <10; i++ {
		wg.Add(1)
		go task1(i, &wg)
	}
	// time.Sleep(1 * time.Second)
	wg.Wait()
}