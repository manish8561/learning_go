I'll provide detailed answers to the key interview questions for a Golang developer position.

Basic Level Questions & Answers:

1. Q: What are the main features of Go that make it unique?
A: Key features of Go include:
- Built-in concurrency support through goroutines and channels
- Fast compilation and efficient garbage collection
- Simplicity and readability of code
- Static typing with type inference
- Cross-platform support
- Built-in testing support
- No inheritance, using composition instead
- Comprehensive standard library

2. Q: Explain the difference between arrays and slices in Go
A: 
```go
// Arrays:
- Fixed length
- Length is part of the type (e.g., [3]int and [4]int are different types)
- When passed to functions, a copy is made
var arr [3]int = [3]int{1, 2, 3}

// Slices:
- Dynamic length
- Built on top of arrays
- Contains reference to underlying array
- Has length and capacity
- More flexible and commonly used
slice := make([]int, 3, 5) // length 3, capacity 5
```

3. Q: What are goroutines and how are they different from threads?
A:
```go
// Goroutines are:
- Lightweight (2KB stack size vs MB for threads)
- Managed by Go runtime instead of OS
- Use channels for communication
- Can run thousands concurrently
- Schedule cooperatively

// Example:
func main() {
    go func() {
        fmt.Println("Running in goroutine")
    }()
    time.Sleep(time.Second)
}
```

4. Q: How do you handle errors in Go?
A:
```go
// Go uses explicit error handling:
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Usage:
result, err := divide(10, 0)
if err != nil {
    log.Fatal(err)
}
```

Intermediate Level Questions & Answers:

1. Q: Explain the difference between `make()` and `new()` in Go
A:
```go
// new():
- Allocates memory but doesn't initialize it
- Returns a pointer to zeroed memory
- Used for value types

ptr := new(int)   // returns *int
*ptr = 42         // dereference to set value

// make():
- Allocates and initializes memory
- Used for slices, maps, and channels
- Returns the type itself, not a pointer

slice := make([]int, 5)    // returns []int
channel := make(chan int)  // returns chan int
```

2. Q: What is the purpose of the `defer` statement?
A:
```go
// defer:
- Delays execution until surrounding function returns
- Executed in LIFO order
- Commonly used for cleanup operations

func processFile() error {
    f, err := os.Open("file.txt")
    if err != nil {
        return err
    }
    defer f.Close()  // Will be called when function returns
    
    // Process file...
    return nil
}
```

3. Q: How do interfaces work in Go?
A:
```go
// Interfaces are implemented implicitly
type Writer interface {
    Write([]byte) (int, error)
}

// Any type that implements Write method satisfies the interface
type FileWriter struct{}

func (fw FileWriter) Write(data []byte) (int, error) {
    // Implementation
    return len(data), nil
}

// Empty interface
interface{} // or any in Go 1.18+
// Can hold values of any type
```

Advanced Level Questions & Answers:

1. Q: Explain Go's context package and its uses
A:
```go
// Context is used for:
- Cancellation signals
- Deadlines
- Passing request-scoped values

func processRequest(ctx context.Context) error {
    // Create a timeout context
    ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
    defer cancel()

    select {
    case <-ctx.Done():
        return ctx.Err()
    case result := <-doWork():
        return processResult(result)
    }
}
```

2. Q: How would you implement a rate limiter?
A:

```go
package main

import (
    "fmt"
    "time"
)

type RateLimiter struct {
    rate  time.Duration
    tokens chan struct{}
    quit   chan struct{}
}

func NewRateLimiter(requestsPerSecond int) *RateLimiter {
    rl := &RateLimiter{
        rate:   time.Second / time.Duration(requestsPerSecond),
        tokens: make(chan struct{}, requestsPerSecond),
        quit:   make(chan struct{}),
    }
    
    // Fill token bucket
    go func() {
        ticker := time.NewTicker(rl.rate)
        defer ticker.Stop()
        
        for {
            select {
            case <-ticker.C:
                select {
                case rl.tokens <- struct{}{}:
                default:
                    // Bucket is full
                }
            case <-rl.quit:
                return
            }
        }
    }()
    
    return rl
}

func (rl *RateLimiter) Allow() bool {
    select {
    case <-rl.tokens:
        return true
    default:
        return false
    }
}

func (rl *RateLimiter) Stop() {
    close(rl.quit)
}

func main() {
    // Create a rate limiter with 2 requests per second
    limiter := NewRateLimiter(2)
    defer limiter.Stop()

    // Simulate requests
    for i := 0; i < 5; i++ {
        if limiter.Allow() {
            fmt.Printf("Request %d allowed\n", i)
        } else {
            fmt.Printf("Request %d throttled\n", i)
        }
        time.Sleep(time.Millisecond * 300)
    }
}

```

3. Q: How would you implement a thread-safe cache with expiration?
A:

```go
package main

import (
    "sync"
    "time"
)

type Cache struct {
    sync.RWMutex
    items map[string]Item
}

type Item struct {
    Value      interface{}
    Expiration int64
}

func NewCache() *Cache {
    cache := &Cache{
        items: make(map[string]Item),
    }
    
    // Start cleanup routine
    go cache.cleanup()
    
    return cache
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
    c.Lock()
    defer c.Unlock()
    
    expiration := time.Now().Add(duration).UnixNano()
    c.items[key] = Item{
        Value:      value,
        Expiration: expiration,
    }
}

func (c *Cache) Get(key string) (interface{}, bool) {
    c.RLock()
    defer c.RUnlock()
    
    item, exists := c.items[key]
    if !exists {
        return nil, false
    }
    
    if time.Now().UnixNano() > item.Expiration {
        return nil, false
    }
    
    return item.Value, true
}

func (c *Cache) cleanup() {
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()
    
    for range ticker.C {
        c.Lock()
        for key, item := range c.items {
            if time.Now().UnixNano() > item.Expiration {
                delete(c.items, key)
            }
        }
        c.Unlock()
    }
}

func main() {
    cache := NewCache()
    
    // Set item with 2 second expiration
    cache.Set("key1", "value1", 2*time.Second)
    
    // Get item
    if value, exists := cache.Get("key1"); exists {
        println(value.(string))
    }
    
    // Wait for expiration
    time.Sleep(3 * time.Second)
    
    // Item should be expired
    if _, exists := cache.Get("key1"); !exists {
        println("Item expired")
    }
}

```

System Design Questions & Answers:

1. Q: How would you design a high-performance REST API in Go?
A: Key points to cover:
- Use of the standard `net/http` package or frameworks like Gin/Echo
- Proper error handling and logging
- Middleware for authentication, logging, and rate limiting
- Connection pooling for database
- Caching strategy
- API versioning
- Input validation
- Documentation (Swagger/OpenAPI)
- Monitoring and metrics (Prometheus)
- Load balancing and scaling considerations


I'll provide more advanced questions and answers that are commonly asked in Golang interviews.

Memory Management & Performance Questions:

1. Q: Explain memory escape analysis in Go
A: 
```go
// Escape analysis determines whether variables should be allocated on stack or heap
// Example of stack allocation:
func createOnStack() int {
    x := 42    // Stays on stack
    return x
}

// Example of heap allocation:
func createOnHeap() *int {
    x := 42    // Escapes to heap
    return &x
}

// Use 'go build -gcflags="-m"' to see escape analysis
```

2. Q: How would you handle memory leaks in Go?
A:
```go
// Common causes and solutions:
1. Goroutine leaks:
func potentialLeak() {
    go func() {
        ch := make(chan int)
        // Channel never read from - goroutine will leak
    }()
}

// Solution: Always ensure proper channel cleanup
func noLeak() {
    ch := make(chan int)
    go func() {
        defer close(ch)
        ch <- 42
    }()
    <-ch
}

2. Large maps with pointers:
// Use sync.Map for concurrent access
// Implement cleanup routines
// Set map to nil when done
```

Advanced Concurrency Questions:

1. Q: Implement a semaphore pattern in Go:
A:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Semaphore struct {
    permits chan struct{}
    wg      sync.WaitGroup
}

func NewSemaphore(maxConcurrent int) *Semaphore {
    return &Semaphore{
        permits: make(chan struct{}, maxConcurrent),
    }
}

func (s *Semaphore) Acquire() {
    s.permits <- struct{}{}
    s.wg.Add(1)
}

func (s *Semaphore) Release() {
    <-s.permits
    s.wg.Done()
}

func (s *Semaphore) Wait() {
    s.wg.Wait()
}

func main() {
    // Create a semaphore allowing 3 concurrent operations
    sem := NewSemaphore(3)
    
    // Launch 5 workers
    for i := 0; i < 5; i++ {
        go func(id int) {
            sem.Acquire()
            defer sem.Release()
            
            // Simulate work
            fmt.Printf("Worker %d starting\n", id)
            time.Sleep(time.Second)
            fmt.Printf("Worker %d done\n", id)
        }(i)
    }
    
    // Wait for all workers to complete
    sem.Wait()
}

```

2. Q: Implement a concurrent pipeline pattern:
A:

```go
package main

import (
    "fmt"
)

func generator(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, n := range nums {
            out <- n
        }
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            out <- n * n
        }
    }()
    return out
}

func filter(in <-chan int, pred func(int) bool) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            if pred(n) {
                out <- n
            }
        }
    }()
    return out
}

func main() {
    // Create pipeline: generate numbers -> square them -> filter even numbers
    c := generator(1, 2, 3, 4, 5)
    out := filter(square(c), func(n int) bool {
        return n%2 == 0
    })
    
    // Consume results
    for n := range out {
        fmt.Println(n)
    }
}

```

3. Q: Implement a fan-out/fan-in pattern:
A:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func producer(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, n := range nums {
            out <- n
        }
    }()
    return out
}

func worker(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            // Simulate work
            time.Sleep(100 * time.Millisecond)
            out <- n * n
        }
    }()
    return out
}

func fanOut(in <-chan int, numWorkers int) []<-chan int {
    workers := make([]<-chan int, numWorkers)
    for i := 0; i < numWorkers; i++ {
        workers[i] = worker(in)
    }
    return workers
}

func fanIn(channels ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    multiplexedStream := make(chan int)

    wg.Add(len(channels))
    for _, c := range channels {
        go func(ch <-chan int) {
            defer wg.Done()
            for i := range ch {
                multiplexedStream <- i
            }
        }(c)
    }

    go func() {
        wg.Wait()
        close(multiplexedStream)
    }()

    return multiplexedStream
}

func main() {
    // Create input stream
    input := producer(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
    
    // Fan out to 3 workers
    workers := fanOut(input, 3)
    
    // Fan in results
    results := fanIn(workers...)
    
    // Collect results
    for result := range results {
        fmt.Println(result)
    }
}

```

Testing & Debugging Questions:

1. Q: How would you write table-driven tests in Go?
A:
```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"zero", 0, 1, 1},
        {"negative numbers", -1, -2, -3},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Add(tt.a, tt.b)
            if got != tt.expected {
                t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expected)
            }
        })
    }
}
```

2. Q: How do you profile Go applications?
A:
```go
// CPU Profiling
import "runtime/pprof"

f, _ := os.Create("cpu.prof")
pprof.StartCPUProfile(f)
defer pprof.StopCPUProfile()

// Memory Profiling
import "runtime/pprof"

f, _ := os.Create("mem.prof")
pprof.WriteHeapProfile(f)

// Use go tool pprof to analyze
// go tool pprof cpu.prof
// go tool pprof mem.prof

// Also use go test -bench=. -benchmem for benchmarking
```

Microservices & API Design:

1. Q: How would you implement circuit breaking in Go?
A:

```go
package main

import (
    "errors"
    "sync"
    "time"
)

type State int

const (
    StateClosed State = iota
    StateHalfOpen
    StateOpen
)

type CircuitBreaker struct {
    mu sync.RWMutex
    
    state                     State
    failures                  int
    failureThreshold         int
    resetTimeout             time.Duration
    lastStateTransitionTime time.Time
}

func NewCircuitBreaker(failureThreshold int, resetTimeout time.Duration) *CircuitBreaker {
    return &CircuitBreaker{
        state:            StateClosed,
        failureThreshold: failureThreshold,
        resetTimeout:     resetTimeout,
    }
}

func (cb *CircuitBreaker) Execute(fn func() error) error {
    if !cb.allowRequest() {
        return errors.New("circuit breaker is open")
    }
    
    err := fn()
    
    cb.mu.Lock()
    defer cb.mu.Unlock()
    
    if err != nil {
        cb.failures++
        if cb.failures >= cb.failureThreshold {
            cb.transitionTo(StateOpen)
        }
        return err
    }
    
    if cb.state == StateHalfOpen {
        cb.transitionTo(StateClosed)
    }
    
    cb.failures = 0
    return nil
}

func (cb *CircuitBreaker) allowRequest() bool {
    cb.mu.RLock()
    defer cb.mu.RUnlock()
    
    switch cb.state {
    case StateClosed:
        return true
    case StateOpen:
        if time.Since(cb.lastStateTransitionTime) > cb.resetTimeout {
            cb.mu.RUnlock()
            cb.mu.Lock()
            cb.transitionTo(StateHalfOpen)
            cb.mu.Unlock()
            cb.mu.RLock()
            return true
        }
        return false
    case StateHalfOpen:
        return true
    default:
        return false
    }
}

func (cb *CircuitBreaker) transitionTo(state State) {
    cb.state = state
    cb.lastStateTransitionTime = time.Now()
}

func main() {
    // Example usage
    cb := NewCircuitBreaker(3, 5*time.Second)
    
    // Simulate some requests
    for i := 0; i < 5; i++ {
        err := cb.Execute(func() error {
            // Simulate failing service
            return errors.New("service error")
        })
        if err != nil {
            println(err.Error())
        }
    }
}

```
