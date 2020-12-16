package main

import "fmt"

var arr = []int{-1, 2, 4, -3, 5, 2, -5, 2}
var n = len(arr)

func main() {
	fmt.Printf("Less Efficient Sum: %v\n", lessEfficientSum())
	fmt.Printf("Efficient Sum: %v\n", efficientSum())
	fmt.Printf("Most Efficient Sum: %v\n", mostEfficientSum())
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func lessEfficientSum() int {
	best := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += arr[k]
			}

			best = Max(best, sum)
		}
	}
	return best
}

func efficientSum() int {
	best := 0
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += arr[j]
			best = Max(best, sum)
		}
	}
	return best
}

func mostEfficientSum() int {
	best := 0
	sum := 0
	for i := 0; i < n; i++ {
		sum = Max(arr[i], sum+arr[i])
		best = Max(best, sum)
	}
	return best
}
