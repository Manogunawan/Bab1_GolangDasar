package main

import "fmt"

func getFullName2() (firstName string, middleName string, lastName string)  {
	firstName = "Manogunawan"
	middleName = "Resqi"
	lastName = "Gultom"
	return
}
func main() {
	a, b, c := getFullName2()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
