package main

import (
	"Tes/helper"
	"fmt"
)

func main()  {
	helper.SayHello("Manogunawan")
	//helper.sayGoodbye("Manogunawan") // error
	fmt.Println(helper.Application)
	//fmt.Println(helper.version) // error
}