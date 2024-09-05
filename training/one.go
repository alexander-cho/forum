package main

import "fmt"

func main() {

	// check
	numItems := 15
	costPerItem := 3.5
	totalCost := costPerItem * float64(numItems)
	fmt.Printf("Spent %.2f on groceries today\n", totalCost)

	// type size
	ageFloat := 24.1
	ageInt := int(ageFloat)
	fmt.Println("You are", ageInt, "years old.")

	// typing
	var username string = "username"
	var password string = "password"
	fmt.Println("Authorization: Basic", username+":"+password)

	//
}
