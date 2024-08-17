package main

import "fmt"

func main() {
    // String 

    var nameOne string = "Abu Khalid"
    var nameTwo = "AbdurRahman"
    var nameThree string

    fmt.Println(nameOne, nameTwo, nameThree)

    nameOne = "Tobiloba"
    nameThree = "Boluwatife"

    fmt.Println(nameOne, nameTwo, nameThree)

    nameFour := "Joseph"

    fmt.Println(nameFour)

    // Integers

    var ageOne int = 26
    var ageTwo = 30
    ageThree := 40

    fmt.Println(ageOne, ageTwo, ageThree)

    // bits && memory

    var numOne int8 = 127
    var numTwo int8 = -128
    var numThree uint8 = 255 

    var scoreOne float32 = 25.98
    var scoreTwo float64 = 12123123131312312312313123.123123
    scoreThree := 1.5 


    fmt.Println(numOne, numTwo, numThree, scoreOne, scoreTwo, scoreThree)
}