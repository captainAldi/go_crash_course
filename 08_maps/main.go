package main

import "fmt"

//Key Value

func main() {
	//Define Maps
	//map[tipeDataKey]tipeDataValue
	email := make(map[string]string)

	//assign kv
	email["bob"] = "bob@gmail.com"
	email["sharon"] = "sharon@gmail.com"
	email["naruto"] = "naruto@gmail.com"

	fmt.Println(email)
	fmt.Println(email["bob"])
	fmt.Println(len(email))

	//Delete
	delete(email, "naruto")

	fmt.Println(email)

	//Create map and Declare
	email2 := map[string]string{"sakura": "sakura@gmail.com", "sasuke": "sasuke@gmail.com"}

	fmt.Println(email2)
}
