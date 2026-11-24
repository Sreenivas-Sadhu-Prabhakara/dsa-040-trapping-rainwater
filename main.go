package main

import "fmt"

// Trap returns how much water the elevation profile holds after rain.
// Two pointers advance from the lower side; the running max on that side
// alone determines the water above each bar.
func Trap(height []int) int {
	l, r := 0, len(height)-1
	leftMax, rightMax, total := 0, 0, 0
	for l < r {
		if height[l] <= height[r] {
			if height[l] >= leftMax {
				leftMax = height[l]
			} else {
				total += leftMax - height[l]
			}
			l++
		} else {
			if height[r] >= rightMax {
				rightMax = height[r]
			} else {
				total += rightMax - height[r]
			}
			r--
		}
	}
	return total
}

func main() {
	fmt.Println(Trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
