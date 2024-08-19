package main

import (
	"fmt"
	"sort"
)

var score = 99.0

func main() {
   sayHello("Joe")
   
   for _, v := range points {
    fmt.Println(v)
   }

   sort.Ints(points)
   fmt.Println(points)
}
   