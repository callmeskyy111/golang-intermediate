package main

import (
	"fmt"
)

type Person struct {
		firstName string
		lastName  string
		age       int
		address Address // embedded struct
		PhHomeCell // anonymous-embedded struct

	}

// Embedded Struct{{}}
type Address struct{
	city string
	country string
}	

// Anonymous struct-fields
type PhHomeCell struct{
	homePh string
	cellPh string
}

// Methods()
func (p Person) fullName()string{
	return p.firstName+" "+p.lastName
}

func (p *Person) incrementAgeByOne(){
	p.age++
}

// main()
func main() {

	p := Person{
		firstName: "Skyy",
		lastName:  "Banerjee",
		age:       29,
		address: Address{
			city: "CCU",
			country: "IN",
		},
		PhHomeCell: PhHomeCell{
			homePh: "+91 9883737478",
			cellPh: "+91 9767866432",
		},
	}

	p1 := Person{
		firstName: "Soumadip",
		// lastName: empty-value
		age: 69,
	}

	p2:= Person{
		firstName: "Soumadip",
		age: 69,
	}

	p1.address.city = "BER"
	p1.address.country = "DE"

	fmt.Println(p.firstName)
	fmt.Println(p1.firstName)
	fmt.Println(p.fullName()) // method
	fmt.Println("Before increment:",p1.age)
	p1.incrementAgeByOne() // method
	fmt.Println("After increment:",p1.age)
	fmt.Println(p.address)
	fmt.Println(p1.address)
	fmt.Println(p.address.city)
	fmt.Println(p1.address.country)
	fmt.Println(p.cellPh)
	fmt.Println(p.address.city)

	//comparison..
	fmt.Println("Are p and p1 equal? -",p==p1)
	fmt.Println("Are p1 and p2 equal? -",p1==p2)

	//Anonymous Structs
	user:=struct {
		userName string
		email string
	}{
		userName: "Akash",
		email: "akash@example.com",
	}

	fmt.Println(user.userName)
}

