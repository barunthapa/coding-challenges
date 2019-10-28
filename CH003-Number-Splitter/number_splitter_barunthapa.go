package main
//package CH003_Number_Splitter

import (
	"fmt"
	"math/rand"
	"time"
)

func split(n, x int) []int {
	var nums []int
	on := n
	sum := 0
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
	for i := x; i > 1; i-- {
		r := r1.Intn(n - x)
		nums = append(nums, r)
		sum = sum + r
		n = n - r
	}
	nums = append(nums, on-sum)
	return nums
}

func main() {
	fmt.Println(split(10, 3))
}
