package main

import "fmt"

func main() {
	//arrays
	var buahArr [2]string

	//assign values
	buahArr[0] = "Apple"
	buahArr[1] = "Jeruk"

	//Deklarasi dan Assign Value disaat bersamaan
	mobilArr := [2]string{"Mercy", "Toyota"}

	//Slice
	//array tanpa specifi index
	// len(slicenya) untuk hitung panjang array nya
	motorArr := []string{"honda", "yamaha"}

	//Print
	fmt.Println(buahArr)
	fmt.Println(buahArr[0])
	fmt.Println(mobilArr)
	fmt.Println(motorArr)
	fmt.Println(len(motorArr))
	//Print range
	//start at 1 stop before 2
	fmt.Println(motorArr[1:2])

}
