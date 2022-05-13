package main

import "fmt"

func main() {
	//for循环的range形式可遍历切片和映射
	var test []int
	test = append(test, 1, 2, 3, 4, 5, 6, 6, 35, 43, 23, 1)
	fmt.Println(test)

	for _, v := range test {
		fmt.Printf("2*%d = %d\n", v, 2*v)
	}
}
