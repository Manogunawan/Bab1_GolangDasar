package main

import "fmt"

func main() {
	name := "Manogunawan"
	counter := 0

	increment := func(){
		name ="Resqi"
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
