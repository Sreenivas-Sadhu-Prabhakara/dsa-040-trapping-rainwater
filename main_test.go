package main

import "testing"

func TestTrap(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
		{[]int{1, 2, 3}, 0},
		{[]int{3, 2, 1}, 0},
		{[]int{}, 0},
		{[]int{5}, 0},
		{[]int{2, 0, 2}, 2},
	}
	for _, c := range cases {
		if got := Trap(c.in); got != c.want {
			t.Errorf("Trap(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
