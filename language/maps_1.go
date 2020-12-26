package main

import "fmt"

func remove(store map[string]int, key string) error { 
   _, ok := store[key] 
   if !ok { 
         return fmt.Errorf("Key not found") 
   } 
   delete(store, key) 
   return nil 
} 

/*
Maps

The Go map is a composite type that is used as containers for storing unordered elements of the same type indexed by 
an arbitrary key value.


In general, map type is specified as follows:

map[<key_type>]<element_type>

The key specifies the type of a value that will be used to index the stored elements of the map. Unlike arrays and 
slices, map keys can be of any type, not just int. Map keys, however, must be of types that are comparable including 
numeric, string, Boolean, pointers, arrays, struct, and interface types.

Similar to a slice, a map value can also be initialized using the make function. Using the make function initializes 
the underlying storage allowing data to be inserted in the map.

Because a map maintains an internal pointer to its backing storage structure, all updates to map parameter within a 
called function will be seen by the caller once the function returns.
*/

func main() {
	var ( 
		legends map[int]string 
		histogram map[string]int 
		calibration map[float64]bool 
		matrix map[[2][2]int]bool    // map with array key type 
		table map[string][]string    // map of string slices 
		// map (with struct key) of map of string 
		log map[struct{name string}]map[string]string 
	)
	var ( 
		histogram map[string]int = map[string]int{ 
		     "Jan":100, "Feb":445, "Mar":514, "Apr":233, 
		     "May":321, "Jun":644, "Jul":113, "Aug":734, 
		     "Sep":553, "Oct":344, "Nov":831, "Dec":312,  
		} 
		table = map[string][]int { 
		     "Men":[]int{32, 55, 12, 55, 42, 53}, 
		     "Women":[]int{44, 42, 23, 41, 65, 44}, 
		} 
	)
	hist := make(map[int]string) 
	hist["Jan"] = 100 
	hist["Feb"] = 445 
	hist["Mar"] = 514
	for key, val := range hist { 
	   adjVal := int(float64(val) * 0.100) 
	   fmt.Printf("%s (%d):", key, val) 
	   for i := 0; i < adjVal; i++ { 
	         fmt.Print(".") 
	   } 
	   fmt.Println() 
	} 
}