package fourth

import (
	"demo/examples"
	"fmt"
	"sync"
)

func Example() {
	defer examples.DisplayLine()

	fmt.Println("Fourth Example handling for concurrently updating a sync map with wait group and go routines")

	var m sync.Map
	var wg sync.WaitGroup

	// Launch multiple goroutines to concurrently updating the map
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			key := fmt.Sprintf("key%d", id)
			value := fmt.Sprintf("value%d", id)

			m.Store(key, value)

		}(i)
	}
	wg.Wait()

	//Read values from the map concurrently
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("key%d", i)
		if val, ok := m.Load(key); ok {
			fmt.Printf("Key: %s, Value: %s \n", key, val)
		}
	}

}
