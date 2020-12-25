package main 

func write(fname string, anagrams map[string][]string) { 
   file, err := os.OpenFile( 
         fname,  
         os.O_WRONLY+os.O_CREATE+os.O_EXCL,  
         0644, 
   ) 
   if err != nil { 
         msg := fmt.Sprintf( 
               "Unable to create output file: %v", err, 
         ) 
         panic(msg) 
   } 
} 
 
func makeAnagrams(words []string, fname string) { 
   defer func() { 
         if r := recover(); r != nil { 
               fmt.Println("Failed to make anagram:", r) 
         } 
   }() 
 
   anagrams := mapWords(words) 
   write(fname, anagrams) 
} 

/*
Function panic recovery

When a function panics, as explained earlier, it can crash an entire program. That may be the desired outcome 
depending on your requirements. It is possible, however, to regain control after a panic sequence has started. 
To do this, Go offers the built-in function called recover. Recover works in tandem with panic. A call to function 
recover returns the value that was passed as an argument to panic.
*/

func main() { 
   words, err := load("") 
   if err != nil { 
         fmt.Println("Unable to load file:", err) 
         os.Exit(1) 
   } 
   makeAnagrams(words, "") 
}