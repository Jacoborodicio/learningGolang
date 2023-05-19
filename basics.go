package basics

import (
  "fmt"
)

type Car struct {
  wheels int
  colour string
  ps int
}

func PrintMessage (message string) {
  fmt.Println(message)
}

func getInfo (c *Car)  (int, string, int) {
  return c.wheels, c.colour, c.ps
}
