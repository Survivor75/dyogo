package main

import (
	"fmt"
)

/*
Channels

When the make function uses the capacity argument, it returns a bidirectional buffered channel.

The buffered channel has the following characteristics:

1. When the channel is empty, the receiver blocks until there is at least one element
2. The sender always succeeds as long as the channel is not at capacity
3. When the channel is at capacity, the sender blocks until at least one element is received

Using a buffered channel, it is possible to send and receive values within the same goroutine without causing a 
deadlock.

*/

func main() {
	ch := make(chan int, 4) // buffered channel

	/*
	The previous code will create a buffered channel with a capacity of 3. The buffered channel operates as a 
	first-in-first-out blocking queue.
	*/

	ch <- 2 
	ch <- 4 
	ch <- 6 
	ch <- 8 

	fmt.Println(<-ch) 
	fmt.Println(<-ch) 
	fmt.Println(<-ch) 
	fmt.Println(<-ch)

}