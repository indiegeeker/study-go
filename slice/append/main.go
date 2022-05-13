package main

import "fmt"

func main() {
	//为切片追加新的元素
	// var test []string
	var test = make([]string, 5)
	//使用append追加新元素
	printSlice(test)

	test = append(test, "zhangan", "lisi") //append之后，切片会按需增长
	printSlice(test)
}

func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
