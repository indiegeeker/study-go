package main

import (
	"fmt"
)

type test struct {
	x int
	y int
}

var (
	//结构体文法通过直接列出字段的值来新分配一个结构体。
	test1      = test{21, 22}
	test_empty = test{}
	//使用 `Name:` 语法可以仅列出部分字段。（字段名的顺序无关。）
	test2 = test{x: 2}
	//特殊的前缀 `&` 返回一个指向结构体的指针。
	p_test = &test{25, 26}
)

func main() {
	// fmt.Println(test{11, 22})

	// //点号访问结构体字段
	// test := test{23, 25}
	// x := test.x
	// y := test.y

	// fmt.Println(x)
	// fmt.Println(y)

	// //结构体指针访问结构体字段
	// p_test := &test
	// p_test.x = 1e9
	// fmt.Println(p_test)

	fmt.Println(test1, test_empty, test2, p_test)
}
