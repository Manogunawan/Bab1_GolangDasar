package main

import "fmt"

type Address struct {
	City, Province, Country string
}

//func ChangeCountryToIndonesia(address Address)
func ChangeCountryToIndonesia(address *Address){
address.Country = "Indonesia"
}

func main() {
	var address1 Address = Address{"Medan", "Sumatera Utara", "Indonesia"}
	//var address4 *Address = &Address{"Medan", "Sumatera Utara", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1
	address2.City = "Jakarta"
	
	*address2 = Address{
		City:     "Padang",
		Province: "Sumatera Barat",
		Country:  "Indonesia",
	}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.City = "Bali"
	fmt.Println(address4)

	var alamat = Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}
	var alamatPointer *Address = &alamat
	//ChangeCountryToIndonesia(&alamat)
	ChangeCountryToIndonesia(alamatPointer)
	fmt.Println(alamat)
}
