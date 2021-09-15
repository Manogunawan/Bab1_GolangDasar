package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
	//fmt.Println("Percobaan Method", man.Name)
}

func main() {
	mano := Man{"Manogunawan"}
	mano.Married()
	fmt.Println(mano.Name)
}
