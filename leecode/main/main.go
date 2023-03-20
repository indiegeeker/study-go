package main

import (
	"fmt"
	"sort"
)

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	reverseNumber := 0
	for x > reverseNumber {
		reverseNumber = reverseNumber*10 + x%10
		fmt.Println(reverseNumber)
		x /= 10
	}
	return x == reverseNumber || x == reverseNumber/10
}

func minimumOperations(nums []int) int {
	//for i := 0; i <= len(nums)-1; i++ {
	//	for j := 0; j < len(nums)-1-i; j++ {
	//		if nums[j] < nums[j+1] {
	//			nums[j], nums[j+1] = nums[j+1], nums[j]
	//		}
	//	}
	//}
	sort.Ints(nums)
	var res, mark int
	for _, num := range nums {
		if num > mark {
			res++
			mark = num
		}
	}
	return res
}

func romanToInt(s string) int {
	/**
	罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
	字符          数值
	I             1
	V             5
	X             10
	L             50
	C             100
	D             500
	M             1000
	*/
	//XIV    XXVII  罗马数字中小的数字在大的数字的右边，累加每个数值
	var reflectValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	n := len(s)
	res := 0
	for index := range s {
		reflectValue := reflectValues[s[index]]
		if index < n-1 && reflectValue < reflectValues[s[index+1]] {
			res -= reflectValue
		} else {
			res += reflectValue
		}
	}
	return res
}

func finalValueAfterOperations(operations []string) int {
	/**
	存在一种仅支持 4 种操作和 1 个变量 X 的编程语言：

	++X 和 X++ 使变量 X 的值 加 1
	--X 和 X-- 使变量 X 的值 减 1
	最初，X 的值是 0

	给你一个字符串数组 operations ，这是由操作组成的一个列表，返回执行所有操作后， X 的 最终值 。
	*/
	res := 0
	for index := range operations {
		if operations[index] == "++X" || operations[index] == "X++" {
			res += 1
		} else {
			res -= 1
		}
	}
	return res
}

func main() {
	//r1 := isPalindrome(1221)
	//r2 := isPalindrome(123421)

	//fmt.Println(r1)
	//fmt.Println(r2)
	nums := []int{1, 5, 0, 3, 5}
	min := minimumOperations(nums)
	fmt.Println(min)

	res := romanToInt("MCMXCIV")
	fmt.Println(res)
}
