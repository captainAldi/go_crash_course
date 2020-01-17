package main

import "fmt"

//Buat Function
//tipe data setelah () jika ingin mereturn sesuatu
func salam(nama string) string {
	return "Hello " + nama
}

//Contoh 2 Parameter
func getSum(num1 int, num2 int) int {
	return num1 + num2
}

//Contoh 3 Parameter
func getSum2(num1, num2 int) int {
	return num1 + num2
}

func main() {
	//karna manggil function maka ata ()
	//masukan parameter nya
	fmt.Println(salam("budi"))
	fmt.Println(getSum(1, 2))
	fmt.Println(getSum2(1, 2))
}
