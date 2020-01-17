package main

import "fmt"

func main() {
	ids := []int{11, 12, 13}

	//Loop Throud ID
	//Range

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//Not Using Index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//Jumlah Semua nya
	sum := 0

	for _, id := range ids {
		sum += id
	}
	fmt.Println("sum", sum)

	//Range with Maps and Key Value
	email2 := map[string]string{"sakura": "sakura@gmail.com", "sasuke": "sasuke@gmail.com"}

	for k, v := range email2 {
		fmt.Printf("%s: %s\n", k, v)
	}

	//just key
	for k := range email2 {
		fmt.Println("Name: " + k)
	}

}
