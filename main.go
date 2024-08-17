package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
    greetings := "Hello there, this is AbdurRahmaan"
    ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

    fmt.Println(strings.Contains(greetings, "AbdurRahmaan"))
    fmt.Println(strings.ReplaceAll(greetings, "Hello", "Hi"))
    fmt.Println(strings.ToUpper(greetings))
    fmt.Println(strings.Split(greetings, " "))
    names := strings.Split(greetings, " ");


    // the original value is unchanged 
    fmt. Println("Original string value =", greetings)


    sort.Ints(ages)
    sort.Strings(names)
    index := sort.SearchInts(ages, 25)
    fmt.Println(ages, index, names)
}
