package main

import (
	"fmt"
	"math"
)

// 整数反转
func reverse(x int) int {
	var y int
	for x != 0 {
		y = y*10 + x%10
		if y < math.MinInt32 || y > math.MaxInt32 {
			return 0
		}
		x = x / 10
	}
	return y
}

func main() {
	fmt.Println(reverse(123))
}
