package main 

import ( 
   "fmt" 
   "math" 
)
 
var ( 
   mul = func(op0, op1 int) int { 
         return op0 * op1 
   } 
   sqr = func(val int) int { 
         return mul(val, val) 
   } 
)

/*
Closures

Go function literals are closures. This means they have lexical visibility to non-local variables declared 
outside of their enclosing code block.

In Go, lexically closed values can remain bounded to their closures long after the outer function that created the
closure has gone out of scope. The garbage collector will handle cleanups as these closed values become unbounded.
*/

func main() { 
   fmt.Printf("mul(25,7) = %d\n", mul(25, 7)) 
   fmt.Printf("sqr(13) = %d\n", sqr(13)) 

   for i := 0.0; i < 360.0; i += 45.0 { 
	rad := func() float64 { 
		return i * math.Pi / 180 
     }()
	fmt.Printf("%.2f Deg = %.2f Rad\n", i, rad) 
  } 
}