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
Function panic

During execution, a function may panic because of any one of following:

1. Explicitly calling the panic built-in function
2. Using a source code package that panics due to an abnormal state
3. Accessing a nil value or an out-of-bound array element
4. Concurrency deadlock

When a function panics, it aborts and executes its deferred calls. Then its caller panics, causing a chain reaction.
*/

func main() { 
   words, err := load("dict.txt") 
   if err != nil { 
         fmt.Println("Unable to load file:", err) 
         os.Exit(1) 
   } 
   anagrams := mapWords(words) 
   write("out.txt", anagrams) 
}