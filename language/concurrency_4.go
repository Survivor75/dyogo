package main

import (
	"fmt"
)

/*
Channels

At declaration, a channel type may also include a unidirectional operator (using the <- arrow again) to indicate 
whether a channel is send-only or receive-only.

<- chan <element_type>: Declares a receive-only channel
chan <- <element_type>: Declares a send-only channel

Once a channel is initialized it is ready for send and receive operations. A channel will remain in that open state 
until it is forcibly closed using the built-in close function.

Once a channel is closed, it has the following properties:

1. Subsequent send operations will cause a program to panic
2. Receive operations never block (regardless of whether buffered or unbuffered)
3. All receive operations return the zero value of the channel's element type


Go offers a long form of the receive operation that returns the value read from the channel followed by a Boolean 
indicating the closed status of the channel. This can be used to properly handle the zero value from a closed 
channel.

*/

func main() { 
	ch := make(chan int, 10) 
	makeEvenNums(4, ch)
	/*
	A bidirectional channel can be converted to a unidirectional channel explicitly or automatically. For instance, 
	when makeEvenNums() is called from main(), it receives the bidirectional channel ch as a parameter. The compiler 
	automatically converts the channel to the appropriate type.
	*/
	for i := 0; i < 4; i++ { 
	 if val, opened := <-ch; opened { 
	       fmt.Println(val) 
	 } else { 
	       fmt.Println("Channel closed!") 
	 } 
	}
	fmt.Println(len(ch)) // current number of elements queued
	fmt.Println(cap(ch)) // declared capacity
} 
 
func makeEvenNums(count int, in chan<- int) { 
   for i := 0; i < count; i++ { 
         in <- 2 * i 
   } 
}