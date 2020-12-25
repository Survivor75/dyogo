package main 
import "fmt" 
 
func div1(op0, op1 int) (int, int) { 
   r := op0 
   q := 0 
   for r >= op1 { 
         q++ 
         r = r - op1 
   } 
   return q, r 
} 

func div2(dvdn, dvsr int) (q, r int) { 
   r = dvdn 
   for r >= dvsr { 
         q++ 
         r = r - dvsr 
   } 
   return 
}

func main() { 
   q, r := div1(71, 5) 
   fmt.Printf("div(71,5) -> q = %d, r = %d\n", q, r) 
}