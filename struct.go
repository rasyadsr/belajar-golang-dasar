package main

import "fmt"

type Customer struct {
	/* biasanya field di awali huruf besar */
	Name, Address string
	Age           int
}

/* struct method */
func (c Customer) sayHi(name string) {
	fmt.Println("hello", name, "my name is", c.Name)
}

func main() {
	/* cara satu */
	var messi Customer
	messi.Name = "Messi"
	messi.Age = 30
	messi.Address = "Argentina"
	messi.sayHi("fulan")

	fmt.Println(messi)
	fmt.Println(messi.Name)
	fmt.Println(messi.Age)
	fmt.Println(messi.Address)

	/* cara dua ( struct literals ) */
	joko := Customer{
		Name:    "joko",
		Address: "Indonesia",
		Age:     18,
	}

	fmt.Println(joko)

	/* cara tiga ( jika menggunakan cara ini akan repot, jika ada perubahan pada field nya seperti posisi )*/
	budi := Customer{
		"Budi",
		"garut",
		19,
	}
	fmt.Println(budi)
}
