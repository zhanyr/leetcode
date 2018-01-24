package main

import "fmt"

func main() {
	oriNums := []int{0, 2,7,10,11,15}
	target := 11

	fmt.Println(twoSum(oriNums, target))

	fmt.Println(testByteMinus('1', '2'))
	fmt.Println(byte(3) + 10)
}

func twoSum(nums []int, target int) []int {
	a := make([]int, 2)
	numMap := make(map[int]int)
	for i,v := range nums {
		if value, ok  := numMap[target-v]; ok{
			if value > i{
				a[0] = i
				a[1] = value
			} else {
				a[0] = value
				a[1] = i
			}
		} else {
			numMap[v] = i
		}

	}
	return a
}

func testByteMinus(a byte, b byte) string {
	return string(int(a) - int(b))
}
