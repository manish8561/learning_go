# Interview Questions

1. What is escape analysis in golang?

   Escape analysis in Go is a compiler technique used to determine whether variables can be safely allocated on the stack or to the heap.

   **Stack vs. Heap Allocation:**

   _Stack Allocation_: Fast memory allocation that occurs within function calls and is automatically cleaned up when the function returns. If a variable’s lifetime ends within the function, it can be allocated on the stack.
   _Heap Allocation_: Slower memory allocation that allows variables to persist beyond the function call, requiring garbage collection for cleanup.

   **How Escape Analysis Works:**

   When you write a Go program, the compiler runs escape analysis to see if a variable "escapes" the function's scope. If the variable is referenced outside its scope (e.g., passed to another goroutine, returned from a function, or assigned to a global variable), it is allocated on the heap.
   If the variable only exists within the function and doesn't escape its scope, it can be safely allocated on the stack, which is more efficient.

2. How memory management is done with goroutines?

   In Go, memory management with goroutines is efficient due to dynamic stack allocation and garbage collection (GC). Goroutines start with a small stack (2 KB) that grows and shrinks as needed, allowing thousands of goroutines to run concurrently. Heap allocation is used for long-lived objects, and Go's concurrent GC handles cleanup without significant performance impacts.

   For concurrency, Go uses channels and synchronization primitives to safely share memory between goroutines. Avoiding goroutine leaks, managing their lifecycles (e.g., using context for cancellation), and careful use of channels are crucial for efficient memory management.

   Garbage collection can be tuned for better performance in high-concurrency environments.

3. What is go service observerablity?

   Observability in Go involves using metrics, logs, and tracing to monitor a service's health and performance in production. By applying these practices, teams can gain deep insights into the system's internal state, enabling faster troubleshooting and more efficient system maintenance.
   e.g.
   Grafarna
   Logging
   Prometheus
   Datadog

4. How garbage collector works in golang?

   Below how it works:

   **Mark-and-Sweep Algorithm**: Go uses a concurrent mark-and-sweep algorithm. During the "mark" phase, the GC identifies live objects (objects still in use) by traversing from root references (global variables, stacks). In the "sweep" phase, it frees the memory occupied by objects that were not marked.

   **Generational Hypothesis**: Go assumes most objects die young, so it focuses on collecting short-lived objects more frequently than long-lived ones.

   **Concurrent Collection**: The Go garbage collector runs concurrently with the application, which minimizes stop-the-world pauses. It pauses the program briefly at the start and end of the garbage collection cycle but runs alongside the program during the mark phase.

   **Tri-color Marking**: The GC uses a tri-color marking system (white, gray, black) to track the state of objects. White objects are unmarked (potentially garbage), gray objects are marked but not fully processed, and black objects are fully processed and known to be in use.

   **Tuning and Optimization**: You can tweak the GC's behavior using environment variables like GOGC (which controls the frequency of collections) or runtime functions to improve performance based on the application's memory usage patterns.
