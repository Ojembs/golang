package main

import (
	"fmt"
)

func main() {
   

    names := [4]string{"joe", "dodo", "pizza", "testing"}

 
    for index, value := range names {
        if index == 1 {
            fmt.Println("Continue at index", index)
            // continue
        } else if index == 2 {
            fmt.Printf("The name of the person at index %v is %v \n", index, value)
        } else {
            fmt.Println("We are done")
        }
    }
}
