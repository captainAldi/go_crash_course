package main

import "fmt"

func main() {
	//Long method

	i := 1

	for i <= 20 {
		fmt.Println(i)
		i++
	}

	//short method
	for i := 1; i <= 10; i++ {
		fmt.Printf("Numbers %d\n", i)
	}

	//Fizz Buzz challange
	//muncul fizz jika bisa dibagi angka 3
	//muncul buzz jika bisa dibagi angka 5
	//muncul fizzbuzz jika bisa dibagi kedua nya (angka 1*angka 2)

	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
