package main

import "fmt"

func printAnyType(val interface{}) { 
   fmt.Println(val) 
}

/*
The empty interface type

The interface{} type, or the empty interface type, is the literal representation of an interface type with an empty 
method set. According to our discussion so far, it can be deduced that all types implement the empty interface since 
all types can have a method set with zero or more members. 

When a variable is assigned the interface{} type, the 
compiler relaxes its build-time type checks. The variable, however, still carries type information that can be 
queried at runtime.

The empty interface is crucially important for idiomatic Go. Delaying type-checking until runtime makes the language 
feels more dynamic without completely sacrificing strong typing. Go offers mechanisms such as type assertion to query 
the type information carried by interfaces at runtime.
*/

func main() { 
   var anyType interface{} 
   anyType = 77.0 
   anyType = "I am a string now" 
   fmt.Println(anyType) 
 
   printAnyType("The car is slow") 
   m := map[string] string{"ID":"12345", "name":"Kerry"} 
   printAnyType(m) 
   printAnyType(1253443455) 
} 
