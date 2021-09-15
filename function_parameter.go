package main

import "fmt"

func sayHelloTo(firstName string, lastName string)  {
	fmt.Println("Hello", firstName, lastName)
}
func main() {
	firstName := "Manogunawan"
	lastName := "Gultom"
	sayHelloTo(firstName,"Manogunawan")
	sayHelloTo(lastName, "Gultom")
}
