package main 
 
import "fmt" 
 
var name, desc string 
var radius int32 
var mass float64
var active bool 
var satellites []string 

func initializedDeclarations() {
	var name, desc string = "Earth", "Planet" 
	var radius int32 = 6378 
	var mass float64 = 5.972E+24 
	var active bool = true 
	var satellites = []string{ 
	  "Moon", 
	}
}


func omittingVariableTypes() {
	var name, desc = "Mars", "Planet" 
	var radius = 6755 
	var mass = 641693000000000.0 
	var active = true 
	var satellites = []string{ 
		"Phobos", 
		"Deimos", 
	}
}

func shortVariableDeclarations() { 
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
  fmt.Println(name) 
  fmt.Println(desc) 
  fmt.Println("Radius (km)", radius) 
  fmt.Println("Mass (kg)", mass) 
  fmt.Println("Satellites", satellites)
}

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