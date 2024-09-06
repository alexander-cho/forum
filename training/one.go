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

	// same line declaration
	averageOpenRate, displayMessage := 0.4554, "is the percentage"
	fmt.Println(averageOpenRate, displayMessage)

	// declaring constants
	const one = "Basic Plan"
	const two = "Premium Plan"
	fmt.Println("plan:", one)
	fmt.Println("plan:", two)

	// computing constants
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour
	fmt.Println("number of seconds in an hour:", secondsInHour)

	// string formatting
	const name = "Alex"
	const numHours = 8.24391
	msg := fmt.Sprintf("I'm %s, and I try to get %.2f hours of sleep a day \n", name, numHours)
	fmt.Print(msg)
}
