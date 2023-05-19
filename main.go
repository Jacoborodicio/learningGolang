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

  // Data Structures
  m := &basics.MaxHeap{}
  fmt.Println(m)
  buildHeap := []int{ 10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
  // Notice how the insert re-arranges the nodes
  for _, v := range buildHeap {
    m.Insert(v)
    fmt.Println(m)
  }

  basics.PrintMessage("---")

  for i := 0; i < 5; i++ {
    m.Extract()
    fmt.Println(m)
  }
}
