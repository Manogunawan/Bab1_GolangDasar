package main

import "fmt"

func Upss(i int) interface{}  {
	if i == 1{
		return 1
	}else if i == 2 {
		return true
	}else {
		return "Upss"
	}
}
func main() {
	var data interface{} = Upss(3)
	fmt.Println(data)
}
