package main

func DivMod(dvdn, dvsr int) (q, r int) { 
  r = dvdn 
  for r >= dvsr { 
    q += 1 
    r = r - dvsr 
  } 
  return 
}

func main() {

	DivMod(75, 5)
}