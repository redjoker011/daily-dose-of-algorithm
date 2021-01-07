// Package main provides ...
package main

import "fmt"

var set = []int{0, 1, 2}
var subset []int
var n = len(set)

func main() {
	for _, v := range set {
		search(v)
	}
}

// Generate Subset recursively
// https://stackoverflow.com/a/56925348/8680724
// search(0);
// -> search(1); // with 0 not in
// ->-> search(2); // with 0 not in AND 1 in
// ->->-> search (3); // with 0 not in AND 1 in AND 2 in. terminates with  (1,2)
// ->->-> search (3); // with 0 not in AND 1 in AND 2 not in. terminates with  (1)
// ->-> search(2); // with 0 not in AND 1 not in
// ->->-> search (3); // with 0 not in AND 1 not in AND 2 in. terminates with  (2)
// ->->-> search (3); // with 0 not in AND 1 not in AND 2 not in. terminates with  ()
// -> search(1); // with 0 in
// ->-> search(2); // with 0 in AND 1 in
// ->->-> search (3); // with 0 in AND 1 in AND 2 in. terminates with (0,1,2)
// ->->-> search (3); // with 0 in AND 1 in AND 2 not in. terminates with (0,1)
// ->-> search(2); // with 0 in AND 1 not in
// ->->-> search (3); // with 0 in AND 1 not in AND 2 in. terminates with  (0,2)
// ->->-> search (3); // with 0 in AND 1 not in AND 2 not in. terminates with  (0)
func search(k int) {
	if k == n {
		// Process Subsets
		if len(subset) >= 1 {
			fmt.Printf("Subset %v\n", subset)
		}
	} else {
		search(k + 1)
		subset = append(subset, k)
		search(k + 1)
		subset = subset[:len(subset)-1]
	}
}
