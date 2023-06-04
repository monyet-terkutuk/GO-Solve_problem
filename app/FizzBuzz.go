package app

import "fmt"

func FizzBuzz() {
	var n int
	fmt.Println("Masukan angka ke -n : ")
	fmt.Scanln(&n)

	for number := 1; number <= n; number++ {
		if number%3 == 0 && number%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if number%3 == 0 {
			fmt.Println("Fizz")
		} else if number%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(number)
		}
	}
}
