package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName)  {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) GetName() string  {
	return person.Name
}

func (animal Animal) GetName() string  {
	return animal.Name
}
func main() {
	var mano Person
	mano.Name ="Manogunawan"

	SayHello(mano)

	cat := Animal{
		Name: "Meong",
	}
	SayHello(cat)
}
