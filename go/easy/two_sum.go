package main

import "fmt"

func twoSum(nums []int, target int) []int {
	result := []int{}

	for i, vOuter := range nums {
		for j, vInner := range nums[i+1:] {
			if vOuter+vInner == target {
				result = append(result, i, j+i+1)
				break
			}
		}

		if len(result) > 0 {
			break
		}
	}

	return result
}

func main() {
	// fmt.Printf("%v %v\n", twoSum([]int{2, 7, 11, 15}, 9), []int{0, 1})
	fmt.Printf("%v %v\n", twoSum([]int{2, 7, 11, 15}, 18), []int{1, 2})
}
