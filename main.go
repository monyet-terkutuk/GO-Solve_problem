package main

import (
	"fmt"
	"go-solve_problem/app"
)

func main() {
	var menu int

	fmt.Println("___________________________________________")
	fmt.Println("___________________ Menu __________________")
	fmt.Println("1. FizzBuzz")
	fmt.Println("2. GuessNumber")
	fmt.Println("___________________________________________\n")
	fmt.Print("Select menu : ")
	fmt.Scan(&menu)

	switch menu {
	case 1:
		app.FizzBuzz()
	case 2:
		app.GuessNumber()
	}

}
