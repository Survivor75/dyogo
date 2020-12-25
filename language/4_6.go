package main 
import ( 
   "fmt" 
   "strings" 
) 


/*
Pointer indirection - accessing referenced values

If all you have is an address, you can access the value to which it points by applying the * operator to the 
pointer value itself (or dereferencing).
*/ 
func main() { 
   a := 3 
   double(&a) 
   fmt.Println(a) 
   p := &struct{ first, last string }{"Max", "Planck"} 
   cap(p) 
   fmt.Println(p) 
} 
 
func double(x *int) { 
   *x = *x * 2 
} 
 
func cap(p *struct{ first, last string }) { 
   p.first = strings.ToUpper(p.first) 
   p.last = strings.ToUpper(p.last) 
}