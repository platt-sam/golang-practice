package main

import (
	"errors"
	"fmt"
)

func getMaxMinFromSlice(nums []int) (max int, min int, err error) {
	if len(nums) == 0 {
		return -1, -1, errors.New("empty slice") // error handling in Go
	}
	var tmpMax, tmpMin = nums[0], nums[0] // assign multiple variables on the same line!
	for i := 1; i < len(nums); i++ {
		if nums[i] > tmpMax {
			tmpMax = nums[i]
		}
		if nums[i] < tmpMin {
			tmpMin = nums[i]
		}
	}
	return tmpMax, tmpMin, nil // nil is Go’s version of null or None
}

func main() {
	numbers := []int{0, 5, 2, 3}
	mx, mn, err := getMaxMinFromSlice(numbers) // calling function and assigning new values
	if err != nil {
		// code to handle error would go here
	}
	fmt.Println(mx, mn)

	mx, _, err = getMaxMinFromSlice(numbers) // just need the maximum? ignore the minimum!
}
