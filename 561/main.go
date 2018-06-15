package main

import (
	"sort"
	"fmt"
)

func main()  {
	nums := []int{1,4,3,2}
	fmt.Println(arrayPairSum(nums))
}

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i,v := range nums{
		if i%2 == 0 {
			sum = sum + v
		}
	}
	return sum
}