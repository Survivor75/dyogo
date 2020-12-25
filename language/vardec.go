package main 
 
import "fmt" 
 
var name, desc string 
var radius int32 
var mass float64
var active bool 
var satellites []string 

/*
Variables
--------------------------------------------------
Variable Declaration Block

var ( 
  name string = "Earth" 
  desc string = "Planet" 
  radius int32 = 6378 
  mass float64 = 5.972E+24
  active bool = true 
  satellites []string   
)

Initialized Declaration

var <identifier list> <type> = <value list or initializer expressions>

var name, desc string = "Earth", "Planet" 
var radius int32 = 6378 
var mass float64 = 5.972E+24 
var active bool = true 
var satellites = []string{ 
  "Moon", 
}

Omitting Variable Types

var <identifier list> = <value list or initializer expressions>

During compilation, the compiler infers the type of the variable based on the assigned value or the initializer 
expression on the right-hand side of the equal sign.

Short Variable Declaration

<identifier list> := <value list or initializer expressions>
name := "Neptune" 
desc := "Planet" 
radius := 24764 
mass := 1.024e26 
active := true 
satellites := []string{ 
     "Naiad", "Thalassa", "Despina", "Galatea", "Larissa", 
 "S/2004 N 1", "Proteus", "Triton", "Nereid", "Halimede", 
     "Sao", "Laomedeia", "Neso", "Psamathe", 
}

For convenience, the short form of the variable declaration does come with several restrictions that you should be 
aware of to avoid confusion:

1. Firstly, it can only be used within a function block
2. The assignment operator :=, declares variable and assign values
3. := cannot be used to update a previously declared variable
4. Updates to variables must be done with an equal sign
*/





/*
Constants
--------------------------------------------------
Typed Constants

const <identifier list> type = <value list or initializer expressions>

const a1, a2 string = "Mastering", "Go" 
const b rune = 'G' 
const c bool = false 
const d int32 = 111009 
const e float32 = 2.71828 
const f float64 = math.Pi * 2.0e+3 
const g complex64 = 5.0i 
const h time.Duration = 4 * time.Second 

Untyped Constants

const <identifier list> = <value list or initializer expression>

const i = "G is" + " for Go " 
const j = 'V' 
const k1, k2 = true, !k1 
const l = 111*100000 + 9 
const m1 = math.Pi / 3.141592 
const m2 = 1.414213562373095048801688724209698078569671
const m3 = m2 * m2 
const m4 = m3 * 1.0e+400 
const n = -5.0i * 3 
const o = time.Millisecond * 5

Constant declaration block

const ( 
  a1, a2 string        = "Mastering", "Go" 
  b      rune          = 'G' 
  c      bool          = false 
  d      int32         = 111009 
  e      float32       = 2.71828 
  f      float64       = math.Pi * 2.0e+3 
  g      complex64     = 5.0i 
  h      time.Duration = 4 * time.Second 
)

Constant Enumeration Block

const ( 
  StarHyperGiant byte = iota 
  StarSuperGiant 
  StarBrightGiant 
  StarGiant 
  StarSubGiant 
  StarDwarf 
  StarSubDwarf 
  StarWhiteDwarf 
  StarRedDwarf 
  StarBrownDwarf 
)

The compiler will then automatically do the following:

1. Declare each member in the block as an untyped integer constant value
2. Initialize iota with a value of zero
3. Assign iota, or zero, to the first constant member (StarHyperGiant)
4. Each subsequent constant is assigned an int value increased by one


You can specify any numeric type that can represent integers or floating point values. For instance, in the 
preceding code sample, each constant will be declared as type byte.

When working with enumerated constants, you may want to throw away certain values that should not be part of 
the enumeration. This can be accomplished by assigning iota to the blank identifier at the desired position 
in the enumeration.
*/

func main() { 
  name = "Sun" 
  desc = "Star" 
  radius = 685800 
  mass = 1.989E+30 
  active = true 
  satellites = []string{ 
    "Mercury", 
    "Venus", 
    "Earth", 
    "Mars", 
    "Jupiter", 
    "Saturn", 
    "Uranus", 
    "Neptune", 
  } 
  fmt.Println(name) 
  fmt.Println(desc) 
  fmt.Println("Radius (km)", radius) 
  fmt.Println("Mass (kg)", mass) 
  fmt.Println("Satellites", satellites) 
}