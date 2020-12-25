package main

import "fmt"

type food interface { 
   eat() 
} 
 
type veggie string 
func (v veggie) eat() { 
   fmt.Println("Eating", v) 
} 
 
type meat string 
func (m meat) eat() { 
   fmt.Println("Eating tasty", m) 
} 
 
func eat(f food) { 
   veg, ok := f.(veggie) 
   if ok { 
         if veg == "okra" { 
               fmt.Println("Yuk! not eating ", veg) 
         }else{ 
               veg.eat() 
         } 
 
         return 
   } 
 
   mt, ok := f.(meat) 
   if ok { 
         if mt == "beef" { 
               fmt.Println("Yuk! not eating ", mt) 
         }else{ 
               mt.eat() 
         } 
         return 
   } 
 
   fmt.Println("Not eating whatever that is: ", f) 
}

/*
Type assertion

When an interface (empty or otherwise) is assigned to a variable, it carries type information that can be queried at 
runtime. Type assertion is a mechanism that is available in Go to idiomatically narrow a variable (of interface type) 
down to a concrete type and value that are stored in the variable.

The general form for type assertion expression is given as follows:

<interface_variable>.(concrete type name)

The expression starts with the variable of the interface type. It is then followed by a dot and the concrete type 
being asserted enclosed in parentheses. The type assertion expression can return two values: one is the concrete 
value (extracted from the interface) and the second is a Boolean indicating the success of the assertion, as shown 
here:

value, boolean := <interface_variable>.(concrete type name)
*/