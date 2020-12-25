package main

import "fmt"


type shape interface { 
   area() float64 
   perim() float64 
} 
 
type rect struct { 
   name string 
   length, height float64 
} 
 
func (r *rect) area() float64 { 
   return r.length * r.height 
} 
 
func (r *rect) perim() float64 { 
   return 2*r.length + 2*r.height 
}
/*
The interface type

The concept of interfaces in Go, similar to other languages, such as Java, is a set of methods that serves as a 
template to describe behavior. A Go interface, however, is a type specified by the interface{} literal, which is 
used to list a set of methods that satisfies the interface.

The interesting aspect of interfaces in Go is how they are implemented and ultimately used. Implementing a Go 
interface is done implicitly. There is no separate element or keyword required to indicate the intent of 
implementation. Any type that defines the method set of an interface type automatically satisfies its implementation.

*/