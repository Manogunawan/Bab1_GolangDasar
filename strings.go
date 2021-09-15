package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Manogunawan", "Mano"))
	fmt.Println(strings.Contains("Resqi", "Nores"))

	//split
	fmt.Println(strings.Split("Manogunawan", " "))

	fmt.Println(strings.ToLower("Manogunawan"))
	fmt.Println(strings.ToUpper("Manogunawan"))

	fmt.Println(strings.ToTitle("manogunawan"))

	fmt.Println(strings.Trim("    Manogunawan    ", " "))
	fmt.Println(strings.ReplaceAll("Mano Joko Mano Mano", "Resqi", "Nores"))
}
