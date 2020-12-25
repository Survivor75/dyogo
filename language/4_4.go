package main 

import "fmt" 
 
var valPtr *float32 
var countPtr *int 
var person *struct { 
   name string 
   age  int 
} 
var matrix *[1024]int 
var row []*int64 

/*
Pointers

In Go, when a piece of data is stored in memory, the value for that data may be accessed directly or a pointer 
may be used to reference the memory address where the data is located. As with other C-family languages, pointers 
in Go provide a level of indirection that let programmers process data more efficiently without having to copy 
the actual data value every time it is needed.

Unlike C, however, the Go runtime maintains control of the management of pointers at runtime. A programmer cannot 
add an arbitrary integer value to the pointer to generate a new pointer address (a practice known as pointer 
arithmetic). Once an area of memory is referenced by a pointer, the data in that area will remain reachable until 
it is no longer referenced any pointer variable. At that point, the unreferenced value becomes eligible for 
garbage collection.

Similar to C/C++, Go uses the * operator to designate a type as a pointer.

Given a variable of type T, Go uses expression *T as its pointer type. The type system considers T and *T as 
distinct and are not fungible. The zero value of a pointer, when it is not pointing to anything, is the address 0, 
represented by the literal constant nil.

Variable aptr, of pointer type *int, is initialized and assigned the address value of variable a using expression &a.

It is also worth noting that Go does not allow the use of the address operator with literal constant for numeric, 
string, and bool types.

There is a syntactical exception to this rule, however, when initializing composite types such as struct and array 
with literal constants.
*/

func main() { 
   fmt.Println(valPtr, countPtr, person, matrix, row) 

   var a int = 1024 
   var aptr *int = &a 

   fmt.Printf("a=%v\n", a) 
   fmt.Printf("aptr=%v\n", aptr)

   structPtr := &struct{ x, y int }{44, 55}
   pairPtr := &[2]string{"A", "B"}
 
   fmt.Printf("struct=%#v, type=%T\n", structPtr, structPtr)
   fmt.Printf("pairPtr=%#v, type=%T\n", pairPtr, pairPtr)

}