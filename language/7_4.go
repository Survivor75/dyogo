package main

import "fmt"

func makePerson() person { 
   addr := address{ 
         city: "Goville", 
         state: "Go", 
         postal: "12345", 
   } 
   return person{ 
         name: "vladimir vivien", 
         address: addr, 
   } 
} 

type person struct { 
   name    string 
   address address 
} 
 
type address struct { 
   street      string 
   city, state string 
   postal      string 
} 

type diameter int 
 
type name struct { 
   long   string 
   short  string 
   symbol rune 
} 
    
type planet struct { 
   diameter 
   name 
   desc string 
}

type Person struct { 
   Name    string `json:"person_name"` 
   Title   string `json:"person_title"` 
   Address `json:"person_address_obj"` 
} 
 
type Address struct { 
   Street string `json:"person_addr_street"` 
   City   string `json:"person_city"` 
   State  string `json:"person_state"` 
   Postal string `json:"person_postal_code"` 
}

/*
Structs

The struct type is constructed by specifying the keyword struct followed by a set of field declarations enclosed 
within curly brackets. In its most common form, a field is a unique identifier with an assigned type which follows 
Go's variable declaration conventions as shown in the previous code snippet (struct also support anonymous fields).

It is crucial to understand that the type definition for a struct includes all of its declared fields.

The names of the embedded types become the field identifiers in the composite literal value for the struct.

To simplify field name resolution, Go follows the following rules when using anonymous fields:

1. The name of the type becomes the name of the field
2. The name of an anonymous field may not clash with other field names
3. Use only the unqualified (omit package) type name of imported types

Fields of an embedded struct can be promoted to its enclosing type. Promoted fields appear in selector expressions 
without the qualified name of their types.

struct variables store actual values. This implies that a new copy of a struct value is created whenever a struct 
variable is reassigned or passed in as a function parameter.

The last topic on structs has to do with field tags. During the definition of a struct type, optional string values 
may be added to each field declaration. The value of the string is arbitrary and it can serve as hints to tools or 
other APIs that use reflection to consume the tags.


*/

func main() {
	var( 
		empty struct{} 
		car struct{make, model string} 
		currency struct{name, country string; code int} 
		node struct{ 
		     edges []string 
		     weight int 
		} 
		person struct{ 
		     name string 
		     address struct{ 
		           street string 
		           city, state string 
		           postal string 
		     } 
		} 
	)

	var( 
	   car = struct{make, model string}{make:"Ford", model:"F150"} 
	   node = struct{ 
	         edges []string 
	         weight int 
	   }{ 
	         edges: []string{"north", "south", "west"}, 
	   } 

	) 

	earth := planet{ 
	     diameter: 7926, 
	     name: name{ 
	           long:   "Earth", 
	           short:  "E", 
	           symbol: '\u2641', 
	     }, 
	     desc: "Third rock from the Sun", 
	}



}