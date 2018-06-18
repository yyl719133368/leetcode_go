package main

import "fmt"

func main() {

	nums := []int{3, 2, 4}
	fmt.Println(twoSum(nums, 6))
}

func twoSum(nums []int, target int) []int {
	l := len(nums)
	result := make([]int, 0)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i)
				result = append(result, j)
				return result
			}
		}
	}
	return result
}
