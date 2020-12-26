package main

import (
	"fmt"
)

func scale(factor float64, vector []float64) []float64 { 
   for i := range vector { 
         vector[i] *= factor 
   } 
   return vector 
} 
 
func contains(val float64, numbers []float64) bool { 
   for _, num := range numbers { 
         if num == val { 
               return true 
         } 
   } 
   return false 
}

func clone(v []float64) (result []float64) { 
   result = make([]float64, len(v), cap(v)) 
   copy(result, v) 
   return 
}

/*
Slices

The slice type is commonly used as the idiomatic construct for indexed data in Go. The slice is more flexible and 
has many more interesting characteristics than arrays. The slice itself is a composite type with semantics similar 
to arrays. In fact, a slice uses an array as its underlying data storage mechanism.

Another way to create a slice value is by slicing an existing array or another slice value (or pointers to these 
values). Go provides an indexing format that makes it easy to express the slicing operation, as follows:

<slice or array value>[<low_index>:<high_index>]

The slicing expression uses the [:] operator to specify the low and high bound indices, separated by a colon, for the slice segment.

1. The low value is the zero-based index where the slice segment starts
3. The high value is the nth element offset where the segment stops

Slicing an existing slice or array value does not create a new underlying array. The new slice creates new pointer 
location to the underlying array.

An array can also be sliced directly. When that is the case, the provided array value becomes the underlying array. 
The capacity and the length the slices will be calculated using the provided array.

Lastly, Go's slice expression supports a longer form where the maximum capacity of the slice is included in the 
expression, as shown here:

<slice_or_array_value>[<low_index>:<high_index>:max]

The max attribute specifies the index value to be used as the maximum capacity of the new slice. That value may be 
less than, or equal to, the actual capacity of the underlying array.

A slice can be initialized at runtime using the built-in function make. This function creates a new slice value and 
initializes its elements with the zero value of the element type. An uninitialized slice has a nil zero value an 
indication that it is not pointing an underlying array. Without an explicitly initialization, with a composite 
literal value or using the make() function, attempts to access elements of a slice will cause a panic.

The make() function takes as an argument the type of the slice to be initialized and an initial size for the slice. Then it returns a slice value. In the previous snippet, make() does the followings:

1. Creates an underlying array of type [6]string
2. Creates the slice value with length and capacity of 6
3. Returns a slice value (not a pointer)

After initialization with the make() function, access to a legal index position will return the zero value for the 
slice element instead of causing a program panic. The make() function can take an optional third parameter that 
specifies the maximum capacity of the slice.

When a function receives a slice as its parameter, the internal pointer of that slice points to the underlying 
array of the slice. Therefore, all updates to the slice, within the function, will be seen by the function's caller.

Recall that assigning or slicing an existing slice value simply creates a new slice value pointing to the same 
underlying array structure. Go offers the copy function, which returns a deep copy of the slice along with a new 
underlying array. The following snippet shows a clone() function, which makes a new copy of a slice of numbers.
*/

func main() {
	var ( 
	    image []byte      
	    ids []string 
	    vector []float64 
	    months []string 
	    q1 []string 
	    histogram []map[string]int
	)
	var ( 
		ids []string = []string{"fe225", "ac144", "3b12c"} 
		vector = []float64{12.4, 44, 126, 2, 11.5}  
		months = []string { 
		     "Jan", "Feb", "Mar", "Apr", 
		     "May", "Jun", "Jul", "Aug", 
		     "Sep", "Oct", "Nov", "Dec", 
		} 
		tables = []map[string][]int { 
		     { 
		           "age":{53, 13, 5, 55, 45, 62, 34, 7}, 
		           "pay":{124, 66, 777, 531, 933, 231}, 
		     }, 
		} 
		graph  = [][][][]int{ 
		     {{{44}, {3, 5}}, {{55, 12, 3}, {22, 4}}}, 
		     {{{22, 12, 9, 19}, {7, 9}}, {{43, 0, 44, 12}, {7}}},     
		} 
	)
	var ( 
	    months [12]string = [12]string{ 
	         "Jan", "Feb", "Mar", "Apr", "May", "Jun", 
	         "Jul", "Aug", "Sep", "Oct", "Nov", "Dec", 
	    } 	 
	    halfyr = months[:6] 
	    q1 = halfyr[:3] 
	    q2 = halfyr[3:6] 
	    q3 = months[6:9] 
	    q4 = months[9:] 
	)
	var ( 
	    months [12]string = [12]string{ 
	         "Jan", "Feb", "Mar", "Apr", "May", "Jun", 
	         "Jul", "Aug", "Sep", "Oct", "Nov", "Dec", 
	    } 
	    summer1 = months[6:9:9] 
	) 
	months := make([]string, 6)
	months := make([]string, 6, 12)
	var vector []float64 
    fmt.Println(len(vector))
    h := make([]float64, 4, 10) 
    fmt.Println(len(h), ",", cap(h))
	months := make([]string, 3, 3) 
   	months = append(months, "Jan", "Feb", "March", "Apr", "May", "June") 
   	months = append(months, []string{"Jul", "Aug", "Sep"}) 
   	months = append(months, "Oct", "Nov", "Dec") 
   	fmt.Println(len(months), cap(months), months)
   	msg := "Bobsayshelloworld!" 
   	fmt.Println( 
         msg[:3], msg[3:7], msg[7:12],  
         msg[12:17], msg[len(msg)-1:], 
   	) 
}