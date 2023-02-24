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
		reflect_value := reflectValues[s[index]]
		if index < n-1 && reflect_value < reflectValues[s[index+1]] {
			res -= reflect_value
		} else {
			res += reflect_value
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

	romanToInt("MCMXCIV")
}
