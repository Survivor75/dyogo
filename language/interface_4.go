package main

type shape interface { 
   area() float64 
} 
 
type polygon interface { 
   shape 
   perim() 
} 
 
type curved interface { 
   shape 
   circonf() 
}

/*
Interface embedding

Another interesting aspects of the interface type is its support for type embedding (similar to the struct type). 
This gives you the flexibility to structure your types in ways that maximize type reuse. Continuing with the shape 
example, the following code snippet reorganizes and reduces the previous interface count from three to two by 
embedding shape into the other two types.

When embedding interface types, the enclosing type will inherit the method set of the embedded types. 
The compiler will complain if the embedded type causes method signatures to clash. Embedding becomes a crucial 
feature, especially when the code applies type validation using type checking. It allows a type to roll up type 
information, thus reducing unnecessary assertion steps.
*/