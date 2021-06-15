package hamming

import "fmt"

// Distance calculates the Hamming distance for two DNA strands
func Distance(a, b string) (int, error) {
	// checks the length of both strands
	if len(a) != len(b) {
		return -1, 	fmt.Errorf("Length of a->%q  and b->%q) is different", a, b)
	}

	var diff = 0
	// compares each strand char by char and calculates the distance
	for i:=0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++		}
	}

	return diff, nil
}
