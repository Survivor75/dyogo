package main 

import "fmt" 

const avogadro float64 = 6.0221413e+23 
const grams = 100.0 

type amu float64 

// function float accepts value of type mass and returns a type of float64 [?]
func (mass amu) float() float64 { 
  return float64(mass) 
} 

type metalloid struct { 
  name   string 
  number int32 
  weight amu 
}

// metalloids is a slice of type metalloid
var metalloids = []metalloid{ 
  metalloid{"Boron", 5, 10.81}, 
  metalloid{"Silicon", 14, 28.085}, 
  metalloid{"Germanium", 32, 74.63}, 
  metalloid{"Arsenic", 33, 74.921}, 
  metalloid{"Antimony", 51, 121.760}, 
  metalloid{"Tellerium", 52, 127.60}, 
  metalloid{"Polonium", 84, 209.0}, 
} 

// function moles accepts value of type mass and returns a type of float64 [?]
func moles(mass amu) float64 { 
  return float64(mass) / grams 
} 

// function atoms accepts value of type float64 and returns a type of float64 [?]
func atoms(moles float64) float64 { 
  return moles * avogadro 
} 

func headers() string { 
  return fmt.Sprintf( 
    "%-10s %-10s %-10s Atoms in %.2f Grams\n", 
    "Element", "Number", "AMU", grams, 
  ) 
} 

func (m metalloid) String() string { 
  return fmt.Sprintf( 
    "%-10s %-10d %-10.3f %e", 
    m.name, m.number, m.weight.float(), atoms(moles(m.weight)), 
  ) 
}

func main() { 
  fmt.Print(headers())
  for _, m := range metalloids { 
    //   fmt.Printf( 
    // "%-10s %-10d %-10.3f %e\n", 
    //   m.name, m.number, m.weight.float(), atoms(moles(m.weight)), 
    //   )
      fmt.Print(m, "\n") 
    }

}