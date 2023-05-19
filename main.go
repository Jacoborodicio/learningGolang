package main

import (
  basics "github.com/Jacoborodicio/learningGolang/mainTopics"
  "fmt"
  // "strconv"
)

func main() {
  // Pointer Basics
  // x := 2
  // basics.PrintSqr(x)
  // basics.PrintMessage(strconv.Itoa(x))
  // basics.AddSqr(&x)
  // basics.PrintMessage("---")
  // basics.PrintMessage(strconv.Itoa(x))

  // Types
  var bmw = basics.Car{Wheels: 4, Colour: "grey", Ps: 184}
  _, c, _ := basics.GetInfo(&bmw)
  fmt.Println(c)
  basics.PrintInfo(&bmw)
}
