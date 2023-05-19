package basics

import (
  "fmt"
)

func Check () {
  fmt.Println("From Pointers")
}

// it prints the value making a local copy of the variable
func PrintSqr(x int) {
  fmt.Println(x*x)
}

// it modifies the original variable
func AddSqr(x *int)  {
 *x *= *x 
}
