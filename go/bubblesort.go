package main

import "fmt"

var arr = [10]int{9, 3, 2, 8, 7, 5, 0, 1, 6, 4}

func main() {
	fmt.Printf("Sorting Collection %v\n", arr)
	for i := len(arr) - 1; i > 1; i-- {
		fmt.Printf("Iterating Outer over %v \n", i)
		for j := 0; j < i; j++ {
			fmt.Printf("Iterating Inner over %v \n", j)
			fmt.Printf("Comparing left %v - %v right %v - %v \n", arr[j], j, arr[j+1], j+1)
			if arr[j] > arr[j+1] {
				fmt.Printf("We have a match %v : %v \n", arr[j], arr[j+1])
				swap(j, j+1)
				fmt.Printf("Collection after swap %v \n", arr)
			}
		}
	}

	fmt.Printf("Collection %v", arr)
}

func swap(left, right int) {
	leftVal := arr[left]
	rightVal := arr[right]
	fmt.Printf("Swapping Values for Left: %v in index %v and Right: %v in index %v \n", leftVal, left, rightVal, right)
	var temp int = arr[left]
	arr[left] = arr[right]
	arr[right] = temp
}
