package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 3, 2, 5}
	target := 5

	Twosum := map[int]int{
		1: 0,
		3: 1,
		2: 2,
		5: 3,
	}

	for i := 0; i < len(arr); i++ {
		ekspetasi := target - arr[i]
		Twosum[arr[i]] = i
		if val, ok := Twosum[ekspetasi]; ok && val != i {
			result := []int{i, val}
			fmt.Println(result)
			// fmt.Println(i)
			return
		}
	}
}

