package main

import (
	"errors"
	"fmt"
)

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

	// concatenation (concat)
	fmt.Println(concat("Hello, ", "my name is Alex"))

	// pass by value (increment)
	x := 5
	increment(x)
	fmt.Println(x)

	// ignoring return values (getName)
	first, _ := getName()
	fmt.Println("Hello there,", first)

	// named return values (yearsUntilEvent)
	fmt.Println(yearsUntilEvent(16))

	// early returns. guard clauses.
	fmt.Println(divide(8, 2))

}

func concat(str1, str2 string) string {
	return str1 + str2
}

func increment(x int) {
	x++
}

func getName() (string, string) {
	return "Alex", "C"
}

func yearsUntilEvent(age int) (
	yearsUntilDriving int,
	yearsUntilDrinking int,
	yearsUntilRental int,
) {
	yearsUntilDriving = 18 - age
	if yearsUntilDriving < 0 {
		yearsUntilDriving = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilRental = 25 - age
	if yearsUntilRental < 0 {
		yearsUntilRental = 0
	}

	// implicit return
	// return

	// explicit return
	return yearsUntilDriving, yearsUntilDrinking, yearsUntilRental

	// // override implicit return
	// return 0, 0, 0

}

func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return dividend / divisor, nil
}
