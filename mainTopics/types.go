package basics

import (
  "fmt"
)

type Car struct {
  Wheels int
  Colour string
  Ps int
}

func PrintMessage (message string) {
  fmt.Println(message)
}

func GetInfo (c *Car)  (int, string, int) {
  return c.Wheels, c.Colour, c.Ps
}

func PrintInfo (c *Car) {
  fmt.Printf("The colour %v with %d Wheels and %d PS", c.Colour, c.Wheels, c.Ps)
}
