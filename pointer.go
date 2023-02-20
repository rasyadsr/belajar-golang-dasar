package main

/**
Ada hal penting yang perlu diketahui mengenai pointer:

1. Variabel biasa bisa diambil nilai pointernya, caranya dengan menambahkan
tanda ampersand ( & ) tepat sebelum nama variabel. Metode ini disebut
dengan referencing.

2. Dan sebaliknya, nilai asli variabel pointer juga bisa diambil, dengan cara
menambahkan tanda asterisk ( * ) tepat sebelum nama variabel. Metode ini
disebut dengan dereferencing.

3. *address2 = Address{"Ubah", "ubah", "ubah"} artinya, merubah referncing
setiap yang merujuk kepada alamat memori yang sama, maka akan berubah datanya

4. keyword new digunakan untuk membuat pointer, namun dengan nilai awal kosong
*/

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}

	// /*
	// 	pada golang, konsep nya pass by value
	// 	jadi, value tersebut di duplicate datanya
	// 	tidak me-reference ke alamat yang sama
	// */
	// address2 := address1
	// address2.City = "Bandung"

	// fmt.Println(address1)
	// fmt.Println(address2)

	// /*
	// 	result :
	// 	{Subang Jawa Barat Indonesia}
	// 	{Bandung Jawa Barat Indonesia}
	// */

	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}

	/*
		operator &, digunakan untuk melakukan pointer,
		agar varibel tersebut mengarah ke memory yang sama
	*/
	address2 := &address1
	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)

	/*
		result :
		{Bandung Jawa Barat Indonesia}
		&{Bandung Jawa Barat Indonesia}
	*/

	*address2 = Address{"Ubah", "ubah", "ubah"}
	fmt.Println(address1)
	fmt.Println(address2)
	/*
		{Ubah ubah ubah}
		&{Ubah ubah ubah}
	*/

	address3 := new(Address)
	fmt.Println(address3)

	address4 := Address{
		City:     "Mars",
		Province: "Alam semesta",
		Country:  "",
	}
	fmt.Println(address4)
	ChangeAddressToIndonesia(&address4)
	fmt.Println(address4)
}
