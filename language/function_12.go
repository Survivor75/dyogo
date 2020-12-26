package main 

import (
   "fmt"
   "os"
   "bufio"
   "errors"
)

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

func load(fname string) ([]string, error) { 
   if fname == "" { 
         return nil, errors.New( 
               "Dictionary file name cannot be empty.")  
   } 
   file, err := os.Open(fname) 
   if err != nil { 
         return nil, fmt.Errorf( 
             "Unable to open file %s: %s", fname, err) 
   } 
   defer file.Close()  
   var lines []string 
   scanner := bufio.NewScanner(file) 
   scanner.Split(bufio.ScanLines) 
   for scanner.Scan() { 
         lines = append(lines, scanner.Text()) 
   } 
   return lines, scanner.Err() 
}

/*
Function panic recovery

When a function panics, as explained earlier, it can crash an entire program. That may be the desired outcome 
depending on your requirements. It is possible, however, to regain control after a panic sequence has started. 
To do this, Go offers the built-in function called recover. Recover works in tandem with panic. A call to function 
recover returns the value that was passed as an argument to panic.

To be able to recover from an unwinding panic sequence, the code must make a deferred call to the recover function. 
In the previous code, this is done in the makeAnagrams function by wrapping recover() inside an anonymous function 
literal.

When the deferred recover function is executed, the program has an opportunity to regain control and prevent the 
panic from crashing the running program.
*/

func main() { 
   words, err := load("") 
   if err != nil { 
         fmt.Println("Unable to load file:", err) 
         os.Exit(1) 
   } 
   makeAnagrams(words, "") 
}