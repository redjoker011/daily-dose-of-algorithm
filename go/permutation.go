// Package main provides ...
package main

import "fmt"

var permutations []int

var set = []int{0, 1, 2}
var n = len(set)
var chosen = make([]bool, 3)

func main() {
	search()
}

func search() {
	if len(permutations) == n {
		// When permutations have same size with set elements means we have a
		// permutation
		fmt.Printf("Permutations %v\n", permutations)
	} else {
		for i := 0; i < n; i++ {
			if chosen[i] {
				continue
			}
			chosen[i] = true
			permutations = append(permutations, i)
			search()
			chosen[i] = false
			permutations = permutations[:len(permutations)-1]
		}
	}
}
