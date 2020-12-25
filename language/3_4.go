package main

import "fmt"

var words = [][]string{  
  {"break", "lake", "go", "right", "strong", "kite", "hello"},  
  {"fix", "river", "stop", "left", "weak", "flight", "bye"},  
  {"fix", "lake", "slow", "middle", "sturdy", "high", "hello"},  
}  
 
func searchBreak(w string) {  
DoSearch:  
  for i := 0; i < len(words); i++ {  
    for k := 0; k < len(words[i]); k++ {  
      if words[i][k] == w {  
        fmt.Println("Found", w)  
        break DoSearch  
      }  
    }  
  }  
}

func searchConntinue(w string) {  
DoSearch:  
  for i := 0; i < len(words); i++ {  
    for k := 0; k < len(words[i]); k++ {  
      if words[i][k] == w {  
        fmt.Println("Found", w)  
        continue DoSearch  
      }  
    }  
  }  
}

/*
The break, continue and goto statements

Here are some attributes of the label for the break statement to remember:

1. The label must be declared within the same running function where the break statement is located.
2. A declared label must be followed immediately by the enclosing control statement (a for loop or switch statement) 
where the break is nested.

Here are some attributes of the label for the continue statement to remember:

1. The label must be declared within the same running function where the continue statement is located
2. The declared label must be followed immediately by an enclosing for loop statement where the continue statement 
is nested.
*/

func main() {

}