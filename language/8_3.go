package main

import "fmt"

type fuel int 

const ( 
    GASOLINE fuel = iota 
    BIO 
    ELECTRIC 
    JET 
) 
type vehicle struct { 
    make string 
    model string 
} 
 
type engine struct { 
   fuel fuel 
   thrust int 
} 
func (e *engine) start() { 
   fmt.Println ("Engine started.") 
} 
 
type truck struct { 
   vehicle 
   engine 
   axels int 
   wheels int 
   class int 
} 
func (t *truck) drive() { 
   fmt.Printf("Truck %s %s, on the go!\n", t.make, t.model)           
} 
 
type plane struct { 
   vehicle 
   engine 
   engineCount int 
   fixedWings bool 
   maxAltitude int 
} 
func (p *plane) fly() { 
   fmt.Printf( 
          "Aircraft %s %s clear for takeoff!\n", 
          p.make, p.model, 
       ) 
}


/*
Objects

Nearly all Go types can play the role of an object by storing states and exposing methods that are capable of 
accessing and modifying those states. The struct type, however, offers all of the features that are traditionally 
attributed to objects in other languages, such as:

1. Ability to host methods
2. Ability to be extended via composition
3. Ability to be sub-typed (with help from the Go interface type)


Go uses the composition over inheritance principle to achieve polymorphism using the type embedding mechanism 
supported by the struct type. In Go, there is no support for polymorphism via type inheritance. Each type is 
independent and is considered to be different from all others. The types truck and plane is a vehicle via a subtype 
relationship.

The struct type embedding mechanism promotes fields and methods when accessed using dot notation. 
For instance, the following fields (make, mode, fuel, and thrust), are all declared in types that are embedded 
inside of the plane type.

They are accessed as if they are members of the plane type when, in fact, they are coming from the types vehicle 
and engine respectively.
*/

func main() { 
   t := &truck { 
         vehicle:vehicle{"Ford", "F750"}, 
         engine:engine{GASOLINE+BIO,700}, 
         axels:2, 
         wheels:6, 
         class:3,     
   } 

   t.start() 
   t.drive() 
 
   p := &plane{} 
   p.make = "HondaJet" 
   p.model = "HA-420" 
   p.fuel = JET 
   p.thrust = 2050 
   p.engineCount = 2 
   p.fixedWings = true 
   p.maxAltitude = 43000 

   p.start() 
   p.fly() 
 
}