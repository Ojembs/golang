package main

import (
	"fmt"
)

func main() {
    // x := 0
    // for x < 5 {
    //     fmt.Println("value of x", x)
    //     x++
    // }

    names := [4]string{"joe", "dodo", "pizza", "testing"}

    // for i := 0; i < len(names); i++ {
    //     fmt.Println(names[i])
    // }

    for index, value := range names {
        fmt.Printf("the value of index %v is %v \n", index, value)
    }
}
