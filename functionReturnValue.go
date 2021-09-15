package main

import "fmt"

func getHello(name string)string  {
	if name == "" {
		return "Hallo Manogunawan"
	} else {
		return "Hello " + name
	}
}
func main() {
	result := getHello("Manogunawan")
	fmt.Println(result)

	fmt.Println(getHello(""))


}
