package main

import "fmt"
type Filter func(string)string

//func sayHelloWithFilter(name string, filter func(string2 string) string){
//	nameFiltered := filter(name)
//	fmt.Println("Hello", nameFiltered)
//}

func sayHelloWithFilter(name string, filter Filter){
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

//func sayHelloWithFilter(name string) {
//	nameFiltered := name
//	if name == "Anjing"{
//		nameFiltered = "..."
//	}
//	fmt.Println("Hello", nameFiltered)
//}

func spanFilter(name string) string{
	if name == "Binatang"{
		return "...."
	}else {
		return name
	}
}
func main() {
	sayHelloWithFilter("Manogunawan", spanFilter)
	sayHelloWithFilter("Binatang", spanFilter)
}
