package main

import "fmt" 

type gallon float64 

func (g gallon) quart() float64 { 
  return float64(g * 4) 
} 

func (g gallon) half() { 
  g = gallon(g * 0.5) 
} 

func (g *gallon) double() { 
  *g = gallon(*g * 2) 
} 

/*
Value and pointer receivers

Similar to regular function parameters, a receiver parameter that uses a pointer to refer to its base value allows 
the code to dereference the original value to update it.

Pointer receiver parameters are widely used in Go. This is because they make it possible to express object-like 
primitives that can carry both state and behaviors.
*/

func main() { 
  var gal gallon = 5 
  gal.half() 
  fmt.Println(gal) 
  gal.double() 
  fmt.Println(gal) 
}