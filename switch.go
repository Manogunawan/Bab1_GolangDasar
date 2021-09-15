package main

import "fmt"

func main() {
	name := "Manogunawan"

	switch name {
	case "Manogunawan":
		fmt.Println("Hai, Manogunawan")
		fmt.Println("Hai, Manogunawan")
	case "Resqi":
		fmt.Println("Hai, Resqi")
		fmt.Println("Hai, Resqi")
	default:
		fmt.Println("Mari berteman")
		fmt.Println("Mari berteman")
	}

	//switch lenght := len(name); lenght > 10 {
	//case true:
	//	fmt.Println("Nama terlalu panjang")
	//case false:
	//	fmt.Println("Nama sudah benar")
	//}

	lenght := len(name)
	switch {
	case lenght > 10:
		fmt.Println("Namamu terlalu panjang")
	case lenght > 5:
		fmt.Println("Namamu lumayan panjang")
	default:
		fmt.Println("Namamu simple ya")
	}
}