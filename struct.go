package main

import "fmt"

type Customer struct {
	Name, Address string
	Age			  int
}

func (customer Customer) sayHi(name string)  {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayHuu(){
	fmt.Println("Huuuu from", a.Name)
}
func main() {
	var mano Customer
	mano.Name = "Nores"
	mano.Address ="Medan"
	mano.Age = 21

	mano.sayHi("Manogunawan")
	mano.sayHuu()
	//sayHi(mano, "Manogunawan")

	fmt.Println(mano.Name)
	fmt.Println(mano.Address)
	fmt.Println(mano.Age)

	resqi := Customer{
		Name:    "Resqi",
		Address: "Jakarta",
		Age:     30,
	}
	fmt.Println(resqi)

	budi := Customer{"Budi", "Lampung", 20}
	fmt.Println(budi)

}
