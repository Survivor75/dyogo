package main

import "fmt"

type truck struct { 
   vehicle 
   engine 
   axels int 
   wheels int 
   class int 
} 

func newTruck(mk, mdl string) *truck { 
   return &truck {vehicle:vehicle{mk, mdl}} 
} 
 
type plane struct { 
   vehicle 
   engine 
   engineCount int 
   fixedWings bool 
   maxAltitude int 
}   

func newPlane(mk, mdl string) *plane { 
   p := &plane{} 
   p.make = mk 
   p.model = mdl 
   return p 
}

/*
The constructor function

Since Go does not support classes, there is no such concept as a constructor. However, one conventional idiom you 
will encounter in Go is the use of a factory function to create and initialize values for a type.

While not required, providing a function to help with the initialization of composite values, such as a struct, 
increases the usability of the code. It provides a place to encapsulate repeatable initialization logic that can 
enforce validation requirements.
*/