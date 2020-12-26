package main

import (
	"fmt"
	"time"
	"math/rand"
)

/*
Arrays

Arrays are static entities that cannot grow or shrink in size once they are declared with a specified length. 
Arrays are a great option when a program needs to allocate a block of sequential memory of a predefined size. 
When a variable of an array type is declared, it is ready to be used without any further allocation semantics.

Similar to C and Java, Go uses the square brackets index expression to access values stored in an array variable. 
This is done by specifying the variable identifier followed by an index of the element enclosed within the square 
brackets.


Arrays values are treated as a single unit. An array variable is not a pointer to a location in memory, but rather 
represents the entire block of memory containing the array elements. This has the implications of creating a new 
copy of an array value when the array variable is reassigned or passed in as a function parameter.

This could have unwanted side effects on memory consumption for a program. One fix for is to use pointer types to 
reference array values.

It should be noted that a composite literal array value can be initialized with the address operator & to initialize 
and return a pointer for the array.
*/

type matrix [2][2][2][2]byte

type numbers [1024 * 1024]int 

func initialize(nums *numbers) { 
   rand.Seed(time.Now().UnixNano()) 
   for i := 0; i < len(nums); i++ { 
         nums[i] = rand.Intn(10000) 
   } 
} 

func max(nums *numbers) int { 
   temp := nums[0] 
   for _, val := range nums { 
         if val > temp { 
               temp = val 
         } 
   } 
   return temp 
}

func initMat() matrix { 
   return matrix{ 
         {{{4, 4}, {3, 5}}, {{55, 12}, {22, 4}}}, 
         {{{2, 2}, {7, 9}}, {{43, 0}, {88, 7}}}, 
   } 
}

func main() {
	// var val [100]int 
	// var days [7]string 
	// var truth [256]bool 
	// var histogram [5]map[string]int 
	// var valInit [100]int = [100]int{44,72,12,55,64,1,4,90,13,54}
	// var daysInit [7]string = [7]string{
	//   "Monday",
	//   "Tuesday",
	//   "Wednesday",
	//   "Thursday",
	//   "Friday",
	//   "Saturday",
	//   "Sunday",
	// }
	// var truthInit = [256]bool{true}
	// var histogramInit = [5]map[string]int {
	//   map[string]int{"A":12,"B":1, "D":15},
	//   map[string]int{"man":1344,"women":844, "children":577},
	// }
	var mat1 matrix 
	mat1 = initMat()
	fmt.Println(mat1)  
	seven := [7]string{"grumpy", "sleepy", "bashful"} 
   	fmt.Println(len(seven), cap(seven))
   	var nums *numbers = new(numbers) 
   	initialize(nums)
}