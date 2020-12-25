package main

import (
	"fmt"
	"math"
)

type shape interface { 
   area() float64 
} 
 
type polygon interface { 
   perim() 
} 
 
type curved interface { 
   circonf() 
} 

type rect struct {...} 

func (r *rect) area() float64 { 
   return r.length * r.height 
} 

func (r *rect) perim() float64 { 
   return 2*r.length + 2*r.height 
} 
 
type triangle struct {...} 

func (t *triangle) area() float64 { 
   return 0.5*(t.a * t.b) 
} 

func (t *triangle) perim() float64 { 
   return t.a + t.b + math.Sqrt((t.a*t.a) + (t.b*t.b)) 
} 
 
type circle struct { ... } 

func (c *circle) area() float64 { 
   return math.Pi * (c.rad*c.rad) 
}

func (c *circle) circonf() float64 { 
   return 2 * math.Pi * c.rad 
}

/*
Implementing multiple interfaces

The implicit mechanism of interfaces allows any named type to satisfy multiple interface types at once. This is 
achieved simply by having the method set of a given type intersect with the methods of each interface type to be 
implemented.

*/