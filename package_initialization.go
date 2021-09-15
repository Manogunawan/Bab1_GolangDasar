package main

import (
	"Tes/database"
	_ "Tes/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
