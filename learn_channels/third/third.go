package third

import (
	"fmt"
	"sync"
	"time"
)

// efficient memory management for extensive memory consumptions

type SomeObject struct {
	Data []byte
}

func createObject() *SomeObject {
	return &SomeObject{
		Data: make([]byte, 1024*1024), // 1 MB data
	}
}

func Example1() {
	// var objects []*SomeObject
	var memoryPiece int
	start := time.Now()

	const worker = 1024 * 1024
	var wg sync.WaitGroup
	// sync.Pool
	// .GET
	// .PUT
	objectPool := sync.Pool{
		New: func() interface{} {
			memoryPiece++
			return createObject()
		},
	}

	//creating large number of objects without resuing them
	for i := 0; i < worker; i++ {
		wg.Add(1)
		// go routines
		go func() {
			defer wg.Done()
			obj := objectPool.Get().(*SomeObject)
			// some use of the objes to do stuff
			// obj := createObject()
			// objects = append(objects, obj)
			objectPool.Put(obj)
		}()
	}
	wg.Wait()

	// release the memory in loop
	// for _, obj := range objects {
	// 	objectPool.Put(obj)
	// }

	// simulate some work
	// time.Sleep(time.Second * 5)

	// the objects are no longer used, but still occupy memory
	end := time.Since(start)
	fmt.Println("Done", memoryPiece)
	fmt.Println("Time2: ", end)
}
