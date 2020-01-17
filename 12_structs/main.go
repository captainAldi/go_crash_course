package main

import (
	"fmt"
	"strconv"
)

//Sama sprti class tapi di go tidak ada class

//Define person struct
type Person struct {
	firstName string
	lastname  string
	city      string
	gender    string
	age       int
}

//method salam (value recevier)
func (p Person) greet() string {
	return "Hello, my name is: " + p.firstName + " " + p.lastname + "And I am " + strconv.Itoa(p.age)
}

//hasBirthDay Method (pointer receiver)
func (p *Person) hasBirthDay() {
	p.age++
}

// getMarried method (poiner receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastname = spouseLastName
	}
}

func main() {
	//init person using struct
	Person1 := Person{firstName: "Samntha", lastname: "SMith", city: "Londong", gender: "f", age: 30}

	//alternative without props name
	Person2 := Person{"wayne", "rooney", "londong", "m", 30}

	fmt.Println(Person1)
	fmt.Println(Person2)

	fmt.Println(Person1.firstName)
	//coba tambah umur
	Person1.age++
	fmt.Println(Person1)

	//coba birhtday tambah umur
	Person1.hasBirthDay()
	fmt.Println(Person1.greet())

	//coba get married
	Person1.getMarried("coba1")
	Person2.getMarried("coba2")

	fmt.Println(Person2.greet())

}
