package main

import "fmt"

func main() {
   age := 26
   name := "Awe AbdurRahman"

   // Print

   fmt.Print("Hello, ")
   fmt.Print("World! \n")
   fmt.Print("Ninja's \n")

   // Println

   fmt.Println("My name is", name, "and I am", age, "years old")

   // printF (Fomatted Strings) %_ = format specifier

   fmt.Printf("My name is %v and I am %v years old \n", name, age)
   fmt.Printf("My name is %q and I am %q years old \n", name, age)
   fmt.Printf("Age is a type of %T \n", age)
   fmt.Printf("Age is a type of %T \n", name)
   fmt.Printf("You scored %f points \n", 225.55)
   fmt.Printf("You scored %0.1f points \n", 225.55)


   // Sprintf (save formatted strings)

   var str = fmt.Sprintf("my name is %v and I am %v years old", name, age)
   fmt.Println("the saved string is:", str)

}