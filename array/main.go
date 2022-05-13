package main

import "fmt"

func main() {
	var num [10]string

	num[0] = "hello"
	num[1] = "world"

	fmt.Println(num)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
