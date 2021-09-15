package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Mano")
	data.PushBack("Manogunawan")
	data.PushBack("Nores")
	data.PushFront("Hai")

	//data.Front().Prev()

	// dari depan ke belakang
	for element := data.Front(); element != nil; element = element.Next(){
		fmt.Println(element.Value)
	}

	for element := data.Back(); element != nil; element = element.Prev(){
		fmt.Println(element.Value)
	}

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)
}
