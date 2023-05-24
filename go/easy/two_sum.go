package main

import (
	"fmt"
	"math/rand"
	"time"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, ok := m[target-n]; ok {
			return []int{j, i}
		}
		m[n] = i
	}
	return nil
}

// using for range
func twoSum2(nums []int, target int) []int {
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
	fmt.Printf("%v %v\n", twoSum([]int{2, 7, 11, 15}, 9), []int{0, 1})
	fmt.Printf("%v %v\n", twoSum([]int{2, 7, 11, 15}, 18), []int{1, 2})

	bigArr := make([]int, 10_000_000)

	for i := 0; i < 10_000_000; i++ {
		bigArr[i] = int(rand.Int63n(int64(1_000_000)))
	}

	start := time.Now()
	fmt.Printf("big map %v %v\n", twoSum(bigArr, 1_500_000), []int{1, 2})
	fmt.Printf("big map %v\n", time.Since(start))

	start = time.Now()
	fmt.Printf("big for-range %v %v\n", twoSum2(bigArr, 1_500_000), []int{1, 2})
	fmt.Printf("big for-range %v\n", time.Since(start))
}
