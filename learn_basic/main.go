package main

import "fmt"

/**
 * Interview questions:
  - What is escape analysis in golang?
    Escape analysis in Go is a compiler technique used to determine whether variables can be safely allocated on the stack or if they need to be moved to the heap.

    Key points:
    Stack vs. Heap Allocation:

    Stack Allocation: Fast memory allocation that occurs within function calls and is automatically cleaned up when the function returns. If a variable’s lifetime ends within the function, it can be allocated on the stack.
    Heap Allocation: Slower memory allocation that allows variables to persist beyond the function call, requiring garbage collection for cleanup.

    How Escape Analysis Works:
    When you write a Go program, the compiler runs escape analysis to see if a variable "escapes" the function's scope. If the variable is referenced outside its scope (e.g., passed to another goroutine, returned from a function, or assigned to a global variable), it is allocated on the heap.
    If the variable only exists within the function and doesn't escape its scope, it can be safely allocated on the stack, which is more efficient.
    Performance Implications:

    Stack allocation is generally faster because memory is automatically freed when the function returns.
    Heap allocation requires more overhead, as the garbage collector must later manage and clean up memory.

  - How memory management is done with goroutines?
    In Go, memory management with goroutines is efficient due to dynamic stack allocation and garbage collection (GC). Goroutines start with a small stack (2 KB) that grows and shrinks as needed, allowing thousands of goroutines to run concurrently. Heap allocation is used for long-lived objects, and Go's concurrent GC handles cleanup without significant performance impacts.

    For concurrency, Go uses channels and synchronization primitives to safely share memory between goroutines. Avoiding goroutine leaks, managing their lifecycles (e.g., using context for cancellation), and careful use of channels are crucial for efficient memory management.

    Garbage collection can be tuned for better performance in high-concurrency environments.

	- What is go service observerablity?
	Observability in Go involves using metrics, logs, and tracing to monitor a service's health and performance in production. By applying these practices, teams can gain deep insights into the system's internal state, enabling faster troubleshooting and more efficient system maintenance.
	Grafarna
	Logging
	Prometheus
*/

/*
*
- How garbage collector works in golang?

Below  how it works:

Mark-and-Sweep Algorithm: Go uses a concurrent mark-and-sweep algorithm. During the "mark" phase, the GC identifies live objects (objects still in use) by traversing from root references (global variables, stacks). In the "sweep" phase, it frees the memory occupied by objects that were not marked.

Generational Hypothesis: Go assumes most objects die young, so it focuses on collecting short-lived objects more frequently than long-lived ones.

Concurrent Collection: The Go garbage collector runs concurrently with the application, which minimizes stop-the-world pauses. It pauses the program briefly at the start and end of the garbage collection cycle but runs alongside the program during the mark phase.

Tri-color Marking: The GC uses a tri-color marking system (white, gray, black) to track the state of objects. White objects are unmarked (potentially garbage), gray objects are marked but not fully processed, and black objects are fully processed and known to be in use.

Tuning and Optimization: You can tweak the GC's behavior using environment variables like GOGC (which controls the frequency of collections) or runtime functions to improve performance based on the application's memory usage patterns.
*/

// defining constants
const (
	first = iota + 2
	second
	third = 'a'
	fourth
	fivth = "mmmmmmm"
	sixth
)

func addLine() {
	fmt.Println("===================================================")
}

// we can also use package reflect.TypeOf or %T formatter in Printf function
// interview question
// type switch example in golang
// also example of polymorphism
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case float32, float64:
		fmt.Printf("Float value:  %v type %T\n", v, v)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	mmap := map[string]string{"0": "a", "1": "b", "2": "c"}

	for key, value := range mmap {
		fmt.Println("Key: ", key, "Value: ", value)
	}

	fmt.Println("const values", first, second, third, fourth, fivth, sixth)

	addLine()

	arr := [...]int{3, 5, 6, 4, 8, 7, 9, 1}

	// bubble sort desecending order
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] < arr[j+1] { // > for ascending
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)

	// insertion sort in ascending order
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println(arr)
	addLine()

	// example to test types with switch
	do(21)
	do(3.14)
	do("hello")
	do(true)
	addLine()

	// new keyword example in golang
	newExample()
}

func newExample() {
	// diff b/w new and make in golang
	/*
	 * new() allocates memory and returns a pointer to the newly allocated zero value of the specified type.
	 * make() is used for initializing slices, maps, and channels, and it returns an initialized (but not zeroed) value of that type.
	 */

	// Using new() to allocate memory for an int
	numPtr := new(int) // numPtr is of type *int, pointing to a zero value (0)

	fmt.Println("Value of numPtr:", *numPtr) // Dereferencing to get the value (initially 0)

	// Updating the value through the pointer
	*numPtr = 42

	fmt.Println("Updated value of numPtr:", *numPtr) // Dereferencing to get the updated value (42)

	// Using new() to allocate memory for a custom struct
	type Person struct {
		name string
		age  int
	}

	personPtr := new(Person) // personPtr is of type *Person, pointing to a zeroed Person struct

	fmt.Println("Person:", *personPtr, "Memory Address: ", &personPtr) // Output: {"" 0}

	// Updating the struct through the pointer
	personPtr.name = "Alice"
	personPtr.age = 30

	fmt.Println("Updated Person:", *personPtr) // Output: {Alice 30}
}
