package main

import "fmt" 
 
type Curr struct { 
  Currency string 
  Name     string 
  Country  string 
  Number   int 
} 
 
var currencies = []Curr{ 
  Curr{"DZD", "Algerian Dinar", "Algeria", 12}, 
  Curr{"AUD", "Australian Dollar", "Australia", 36}, 
  Curr{"EUR", "Euro", "Belgium", 978}, 
  Curr{"CLP", "Chilean Peso", "Chile", 152}, 
  Curr{"EUR", "Euro", "Greece", 978}, 
  Curr{"HTG", "Gourde", "Haiti", 332},
} 
 
func isDollar(curr Curr) bool { 
  var bool result 
  switch curr { 
  default: 
    result = false 
  case Curr{"AUD", "Australian Dollar", "Australia", 36}: 
    result = true 
  case Curr{"HKD", "Hong Kong Dollar", "Hong Koong", 344}: 
    result = true 
  case Curr{"USD", "US Dollar", "United States", 840}: 
    result = true 
  } 
  return result 
} 
func isDollar2(curr Curr) bool { 
  dollars := []Curr{currencies[2], currencies[6], currencies[9]} 
  switch curr { 
  default: 
    return false 
  case dollars[0]: 
    fallthrough 
  case dollars[1]: 
    fallthrough 
  case dollars[2]: 
    return true 
  } 
  return false 
} 
 
func isEuro(curr Curr) bool { 
  switch curr { 
  case currencies[2], currencies[4], currencies[10]: 
    return true 
  default: 
    return false 
  } 
} 
/*
SWITCH Statements

The expressions in case clauses are evaluated from left to right, top to bottom, until a value (or expression) is 
found that is equal to that of the switch expression. Upon encountering the first case that matches the switch 
expression, the program will execute the statements for the case block and then immediately exit the switch block. 
Unlike other languages, the Go case statement does not need to use a break to avoid falling through the next case.

The switch keyword may be immediately followed by a simple initialization statement where variables, 
local to the switch code block, may be declared and initialized.

Given Go's strong type support, it should be of little surprise that the language supports the ability to query 
type information. The type switch is a statement that uses the Go interface type to compare the underlying type 
information of values (or expressions).

*/ 
func main() { 
  curr := Curr{"EUR", "Euro", "Italy", 978} 
  if isDollar(curr) { 
    fmt.Printf("%+v is Dollar currency\n", curr) 
  } else if isEuro(curr) { 
    fmt.Printf("%+v is Euro currency\n", curr) 
  } else { 
    fmt.Println("Currency is not Dollar or Euro") 
  } 
  dol := Curr{"HKD", "Hong Kong Dollar", "Hong Koong", 344} 
  if isDollar2(dol) { 
    fmt.Println("Dollar currency found:", dol) 
  } 
}