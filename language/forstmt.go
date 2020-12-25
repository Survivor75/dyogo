package main

import (
  "fmt"
  "math/rand"
)

type Curr struct {  
  Currency string  
  Name     string  
  Country  string  
  Number   int  
}

var currencies = []Curr{  
  Curr{"KES", "Kenyan Shilling", "Kenya", 404},  
  Curr{"AUD", "Australian Dollar", "Australia", 36}} 

var list1 = []string{ 
  "break", "lake", "go",  
  "right", "strong",  
  "kite", "hello",
}  
 
var list2 = []string{ 
  "fix", "river", "stop",  
  "left", "weak", "flight",  
  "bye",
}

func listCurrs(howlong int) {  
  i := 0  
  for i < len(currencies) {  
    fmt.Println(currencies[i])  
    i++  
  }  
}

func sortByNumber() {  
  N := len(currencies)  
  for i := 0; i < N-1; i++ {  
     currMin := i  
     for k := i + 1; k < N; k++ {  
      if currencies[k].Number < currencies[currMin].Number {  
         currMin = k  
    }  
     }  
     if currMin != i {  
        temp := currencies[i]  
    currencies[i] = currencies[currMin]  
    currencies[currMin] = temp  
   } 
  }  
}

func nextPair() (w1, w2 string) {  
  pos := rand.Intn(len(list1))  
  return list1[pos], list2[pos]  
}  
 
func main() {  
  rand.Seed(31)  
  for w1, w2:= nextPair();  w1 != "go" && w2 != "stop";  w1, w2 = nextPair() {  
    fmt.Printf("Word Pair -> [%s, %s]\n", w1, w2)  
  }  

  vals := []int{4, 2, 6} 

  for i, v := range vals { 
    vals[i] = v - 1 
  } 

  fmt.Println(vals) 

  for i := range currencies { 
    fmt.Printf("%d: %v\n", i, currencies[i])
  }
}