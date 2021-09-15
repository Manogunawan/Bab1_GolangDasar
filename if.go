package main

import "fmt"

func main() {
	var name = "Nores"

	if name == "Mano"{
		fmt.Println("Hello Mano")
	}else if name == "Nores"{
		fmt.Println("Hello Nores")
	} else {
		fmt.Println("Hi, cepu")
	}


	if length := len(name) ;length> 5 {
		fmt.Println("Panjang sekali")
	}else {
		fmt.Println("Nama Sudah ")
	}
}
