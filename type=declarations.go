package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpMano NoKTP = "127188881819"
	var marriedStatus Married = true
	fmt.Println(noKtpMano)
	fmt.Println(marriedStatus)
}
