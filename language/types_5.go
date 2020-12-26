package main 

import "fmt" 

/*
The new() function

The built-in function new(<type>) can also be used to initialize a pointer value. It first allocates the 
appropriate memory for a zero-value of the specified type. The function then returns the address for the newly 
created value. The following program uses the new() function to initialize variables intptr and p.

However, when dealing with composites, the idiom is more forgiving. It is not necessary to write *p.first to access 
the pointer's field value. We can drop the * and just use p.first = strings.ToUpper(p.first).
*/

func main() { 
   intptr := new(int) 
   *intptr = 44 
   p := new(struct{ first, last string }) 
   p.first = "Samuel" 
   p.last = "Pierre" 
   fmt.Printf("Value %d, type %T\n", *intptr, intptr) 
   fmt.Printf("Person %+v\n", p) 
}