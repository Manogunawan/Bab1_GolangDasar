package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("m([a-z])o")

	fmt.Println(regex.MatchString("mno"))
	fmt.Println(regex.MatchString("mxo"))

	search := regex.FindAllString("mno mxo eki", 2)
	fmt.Println(search)
}
