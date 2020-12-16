// Binary search only works with ascending sorted array
package main

import "fmt"

var arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {
	lessEfficient()
	efficient()
}

// Halve the array size in half for every iteration
func lessEfficient() {
	var n = 7

	var a = 0
	var b = len(arr) - 1
	for a <= b {
		var m = (a + b) / 2

		if m > n {
			b = m - 1
		} else if m < n {
			a = m + 1
		} else {
			fmt.Printf("N %v found in index %v\n", n, m)
			return
		}
	}
}

// Jump through array and halve the length of jumps until we reach n==1
func efficient() {
	x := 8
	k := 0

	for i := len(arr) / 2; i >= 1; i /= 2 {
		for k+i < len(arr) && arr[k+i] <= x {
			k += i
		}
	}

	if arr[k] == x {
		fmt.Printf("N %v found in index %v\n", x, k)
	}
}
