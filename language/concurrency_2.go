package main

import (
	"fmt"
)

/*
Channels

When talking about concurrency, one of the natural concerns that arises is that of data safety and synchronization 
among concurrently executing code. If you have done concurrent programming in languages such as Java or C/C++, you 
are likely familiar with the, sometimes brittle, choreography required to ensure running threads can safely access 
shared memory values to achieve communication and synchronization between threads.

This is one area where Go diverges from its C lineage. Instead of having concurrent code communicate by using shared 
memory locations, Go uses channels as a conduit between running goroutines to communicate and share data.

Do not communicate by sharing memory; instead, share memory by communicating.

The concept of channel has its roots in communicating sequential processes (CSP), work done by renowned computer 
scientist C. A. Hoare, to model concurrency using communication primitives. Channels provide the means to 
synchronize and safely communicate data between running goroutines.

The channel type declares a conduit within which only values of a given element type may be sent or received by the 
channel. The chan keyword is used to specify a channel type, as shown in the following declaration format:

chan <element type>

Go uses the <- (arrow) operator to indicate data movement within a channel.



When the arrow is placed to the left of the value, variable or expression, it indicates a send operation to the 
channel it points to. 
Example: intCh <- 12
In this example, 12 is sent into channel intCh.

When the <- operator is place to the left of a channel, it indicates a receive operation from the channel. The value 
variable is assigned the value received from the intCh channel.

An uninitialized channel has a nil zero value and must be initialized using the built-in make function. 

A channel can be initialized as either unbuffered or buffered, depending on its specified capacity. Each of type of 
channel has different characteristics that are leveraged in different concurrency constructs.

When the make function is invoked without the capacity argument, it returns a bidirectional unbuffered channel.

1. If the channel is empty, the receiver blocks until there is data
2. The sender can send only to an empty channel and blocks until the next receive operation
3. When the channel has data, the receiver can proceed to receive the data.
*/

func main() {

	var ch1 chan int
	ch2 := make(chan int)
	
	/*
	Sending to an unbuffered channel can easily cause a deadlock if the operation is not wrapped in a goroutine.
	ch2 <- 12 // blocks
	The sender blocks immediately upon sending to an unbuffered channel. This means any subsequent statement, to 
	receive from the channel for instance, becomes unreachable, causing a deadlock.
	fmt.Println(<-ch2)
	*/

	go func() { ch <- 12 }()

	/*
	The send operation is wrapped in an anonymous function invoked as a separate goroutine. This allows the main 
	function to reach the receive operation without blocking.
	*/

	fmt.Println(<-ch2)
}