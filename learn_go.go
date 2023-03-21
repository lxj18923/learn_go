package main

import "fmt"

func main() {
	nums := []int{3, 5, 1, 6, 8, 4, 2}
	fmt.Println("Before sorting:", nums)
	quickSort(nums)
	fmt.Println("After sorting:", nums)
}

func quickSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	pivot := nums[0]
	left, right := 1, len(nums)-1
	for left <= right {
		if nums[left] <= pivot {
			left++
		} else {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		}
	}
	nums[0], nums[right] = nums[right], nums[0]
	quickSort(nums[:right])
	quickSort(nums[right+1:])
	// 我是修改代码的
}
