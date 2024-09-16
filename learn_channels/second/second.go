package second

import (
	"fmt"
	"sync"
	"time"
	"demo/examples"
	"sync/atomic"

)

// synchronization with mutex and atomic
func Example1() {
	defer examples.DisplayLine()
	fmt.Println("Example 1 from second for synchornization")

	var counter int32
	var wg sync.WaitGroup
	start := time.Now()
	var mu sync.Mutex

	for i := 0; i < 10000000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			counter++
			atomic.AddInt32(&counter, 1)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Since(start)
	fmt.Println("Counter: ", counter)
	fmt.Println("Atomic counter state: ", atomic.LoadInt32(&counter))
	fmt.Println("Time: ", end)
}
