package main

import (
	"fmt"
	"sync"
)

type counter struct {
	sync.Mutex
	counter int
}

func (c *counter) Increment() {
	c.Lock()
	defer c.Unlock()
	c.counter++
}

func main() {
	c := new(counter)
	var wg sync.WaitGroup
	for i := 0; i < 111; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Increment()
		}()
	}
	wg.Wait()
	fmt.Println(c.counter)
}
