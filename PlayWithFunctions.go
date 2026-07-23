package main

import (
	"fmt"
)

//functions with multiple return value giving us the minimum and maximum
func minMax(nums ...int) (minimum int, maximum int) {
	minimum = nums[0]
	maximum = nums[0]
	for _, val := range nums {
		if val < minimum {
			minimum = val
		}
		if val > maximum {
			maximum = val
		}
	}
	return
}

func functionClosurePlay() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {

	fmt.Println("welcome to the functions Demo...!")
	min, max := minMax(2, 3, 4, 5, 7, 88, 3, 4, 66, 99, 2)
	fmt.Println("The minimum is", min, " and the maximum is", max)

	fmt.Println("-----------------------------------------")
	closureVar := functionClosurePlay()

	fmt.Println(closureVar())
	fmt.Println(closureVar())
	fmt.Println(closureVar())

}
