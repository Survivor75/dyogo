package main 
import "fmt" 
 
func do(steps ...string) { 
   defer fmt.Println("All done!") 
   for _, s := range steps { 
         defer fmt.Println(s) 
   } 
 
   fmt.Println("Starting") 
} 

/*
Deferring function calls

Go supports the notion of deferring a function call. Placing the keyword defer before a function call has the 
interesting effect of pushing the function unto an internal stack, delaying its execution right before the enclosing
function returns. To better explain this, let us start with the following simple program that illustrates the use of
defer.

The defer keyword modifies the execution flow of a program by delaying function calls. One idiomatic usage for this 
feature is to do a resource cleanup. Since defer will always get executed when the surrounding function returns, 
it is a good place to attach cleanup code such as:

1. Closing open files
2. Releasing network resources
3. Closing the Go channel
4. Committing database transactions
5. And do on
*/

func main() { 
   do( 
         "Find key", 
         "Aplly break", 
         "Put key in ignition", 
         "Start car", 
   ) 
}