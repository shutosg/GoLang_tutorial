package main

import "fmt"

func main() {
	fmt.Println("Hello, playground")
	for i := 1; i<101; i++ {
		fmt.Println(fizzbuzz(i))
	}
}

func fizzbuzz(i int) string{
	if i%15 == 0 {
		return "Fizz Buzz"
	} else if i%5 == 0 {
		return "Buzz"
	} else if i%3 == 0 {
		return "Fizz"
	} else {
		return string(i)
	}
	
}
