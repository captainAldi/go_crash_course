package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	//Using var
	//Boleh untuk tidak menuliskan tipe data karna akan otomatis
	var name string = "Rooney"
	var age int = 30
	var isCool bool = true

	//Shorthand untuk deklarasi variable perhatikan scope
	nama2 := "Wesley"
	size := 1.5

	//One Line Deklarasi
	email, kota := "tes@gmail.com", "jambi"

	//specific boleh
	var hp int32 = 2

	fmt.Println(name, age)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", hp)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", nama2)
	fmt.Printf("%T\n", size)
	fmt.Printf("%T\n", email)
	fmt.Printf("%T\n", kota)
}
