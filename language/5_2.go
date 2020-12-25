package main 

import "fmt" 
 
func avg(nums ...float64) float64 { 
   n := len(nums) 
   t := 0.0 
   for _, v := range nums { 
         t += v 
   } 
   return t / float64(n) 
} 
 
func sum(nums ...float64) float64 { 
   var sum float64 
   for _, v := range nums { 
         sum += v 
   } 
   return sum 
} 

/*
Variadic parameters

The last parameter of a function can be declared as variadic (variable length arguments) by affixing ellipses (…) 
before the parameter's type. This indicates that zero or more values of that type may be passed to the function 
when it is called.

The compiler resolves the variadic parameter as a slice of type []float64 in both the preceding functions. 
The parameter values can then be accessed using a slice expression.

When no parameters are provided, the function receives an empty slice. The astute reader may be wondering, 
"Is it possible to pass in an existing slice of values as variadic arguments?" Thankfully, Go provides an easy 
idiom to handle such a case. The slice can be passed as a variadic parameter by adding ellipses to the parameter 
in the sum(points...) function call.

*/ 
func main() { 
   fmt.Printf("avg([1, 2.5, 3.75]) =%.2f\n", avg(1, 2.5, 3.75)) 
   points := []float64{9, 4, 3.7, 7.1, 7.9, 9.2, 10} 
   fmt.Printf("sum(%v) = %.2f\n", points, sum(points...)) 
}