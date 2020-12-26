package main

import (
	"fmt"
)

/*
Goroutines

Go has its own concurrency primitive called the goroutine, which allows a program to launch a function (routine) to 
execute independently from its calling function. Goroutines are lightweight execution contexts that are multiplexed 
among a small number of OS-backed threads and scheduled by Go's runtime scheduler. That makes them cheap to create 
without the overhead requirements of true kernel threads. As such, a Go program can initiate thousands (even hundreds
of thousands) of goroutines with minimal impact on performance and resource degradation.

Goroutines are launched using the go statement as follows:

go <function or expression>

A goroutine is created with the go keyword followed by the function to schedule for execution. The specified 
function can be an existing function, an anonymous function, or an expression that calls a function.

When using the go statement with a function literal, it is treated as a regular closure with lexical access to 
non-local variables.

This is safe as long as the variables captured in the closure are not expected to change after the goroutine starts. 
If these values are updated outside of the closure, it may create race conditions causing the goroutine to read 
unexpected values by the time it is scheduled to run.

In general, all goroutines run independently of each other. A function that creates a goroutine does not wait for it 
to return, it continues with its own execution stream unless there is a blocking condition.

Go's runtime scheduler uses a form of cooperative scheduling to schedule goroutines. By default, the scheduler will 
allow a running goroutine to execute to completion. However, the scheduler will automatically yield to another 
goroutine for execution if one of the following events occurs:

1. A go statement is encountered in the executing goroutine
2. A channel operation is encountered
3. A blocking system call (file or network IO for instance) is encountered
4. After the completion of a garbage collection cycle

The scheduler will schedule a queued goroutines ready to enter execution when one of the previous events is 
encountered in a running goroutine. It is important to point out that the scheduler makes no guarantee of the order 
of execution of goroutines.
*/

func count(start, stop, delta int) { 
   for i := start; i <= stop; i += delta { 
         fmt.Println(i) 
   } 
} 

func main() { 
	go count(10, 50, 10) 
	go count(60, 100, 10) 
	go count(110, 200, 20) 
	starts := []int{10,40,70,100} 

	/*
	The goroutine closures, invoked with each loop iteration, receive a copy of the j variable via the function 
	parameter. This creates a local copy of the j value with the proper value to be used when the goroutine is 
	scheduled to run.
	*/
	
	for _, j := range starts{ 
	     go func() { 
	           count(j, j+20, 10) 
	     }() 
	}
	fmt.Scanln() // blocks for kb input
} 