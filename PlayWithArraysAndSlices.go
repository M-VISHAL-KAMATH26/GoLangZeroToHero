package main

import "fmt"

//function to find the average where the parameters is an slice of integer and return value will be float64
func findAvg(nums []int) float64 {
	var avg float64 = 0
	var totalSum int = 0
	for _, val := range nums {
		totalSum += val
		avg = float64(totalSum) / float64(len(nums))
	}
	return avg
}

//function that takes the slice and doubles things in place on same slice
func doubleSliceElementsInPlace(nums []int) {
	for i, _ := range nums {
		nums[i] *= 2
	}
}

//function that takes the slice and gives a new slice of doubled value of each elements without modifying the original one
func doubleSliceElementsReturnNew(nums []int) []int {
	resultantSlice := make([]int, len(nums))
	for i, _ := range nums {
		resultantSlice[i] = nums[i] * 2
	}
	return resultantSlice
}

func main() {
	fmt.Println("welcome to testing arrays and slices")

	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	fmt.Println(arr)

	var arr2 = arr
	arr2[3] = 777
	fmt.Println(arr2)
	fmt.Println(arr)

	var slices []int = []int{11, 22, 33, 44, 55}
	var slice2 = slices[:]
	slice2[0] = 9999
	fmt.Println(slices)
	fmt.Println(slice2)

	fmt.Println("---------------------------------------")
	slice3 := []string{"apple", "banana", "cherry"}
	fmt.Println(slice3, " is the slice and the length of the slice is  ", len(slice3), " and the capacity of the slice is ", cap(slice3))
	slice3 = append(slice3, "date")
	fmt.Println(slice3, " is the slice and the length of the slice is  ", len(slice3), " and the capacity of the slice is ", cap(slice3))
	fmt.Println(slice3[0:2])

	fmt.Println("---------------------------------------")
	fmt.Println(findAvg([]int{10, 10, 10, 10, 10}))

	fmt.Println("---------------------------------------")
	fmt.Println(slice2)
	doubleSliceElementsInPlace(slice2)
	fmt.Println(slice2)

	var slice4 []int = []int{10, 20, 30, 40, 50}
	fmt.Println(slice4)
	slice4DoubledRes := doubleSliceElementsReturnNew(slice4)
	fmt.Println(slice4)
	fmt.Println(slice4DoubledRes)

}
