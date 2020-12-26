package main

import "fmt"

type ounce float64 

func (o ounce) cup() cup { 
   return cup(o * 0.1250) 
} 
 
type cup float64 

func (c cup) quart() quart { 
   return quart(c * 0.25) 
} 

func (c cup) ounce() ounce { 
   return ounce(c * 8.0) 
} 
 
type quart float64 

func (q quart) gallon() gallon { 
   return gallon(q * 0.25) 
} 

func (q quart) cup() cup { 
   return cup(q * 4.0) 
} 
 
type gallon float64 

func (g gallon) quart() quart { 
   return quart(g * 4) 
} 

/*
Go Methods

A Go function can be defined with a scope narrowed to that of a specific type. When a function is scoped to a type, 
or attached to the type, it is known as a method. A method is defined just like any other Go function. However, its 
definition includes a method receiver, which is an extra parameter placed before the method's name, used to specify 
the host type to which the method is attached.


The number of methods attached to a type, via the receiver parameter, is known as the type's method set. 
This includes both concrete and pointer value receivers. The concept of a method set is important in determining 
type equality, interface implementation, and support of the notion of the empty method set for the empty interface.

It is important to note that the base type for method receivers cannot be a pointer (nor an interface).
*/

func main() { 
    gal := gallon(5) 
    fmt.Printf("%.2f gallons = %.2f quarts\n", gal, gal.quart()) 
    ozs := gal.quart().cup().ounce() 
    fmt.Printf("%.2f gallons = %.2f ounces\n", gal, ozs) 
}