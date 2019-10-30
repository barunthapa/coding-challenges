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
	for i := x; i > 1; i-- {
		if n <= x {
			nums = append(nums, 1)
			continue
		}
		r := getRand(n - x)
		nums = append(nums, r)
		sum = sum + r
		n = n - r
	}
	fmt.Println("Sum:", sum)
	nums = append(nums, on-sum)
	return nums
}

func main() {
	fmt.Println(split(100, 5))
}

func getRand(num int) int {
	fmt.Println("num", num)
	if num < 1 {
		return 0
	}
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	n := r.Intn(num)
	fmt.Println(n)
	for ; n == 0 ; {
		n = r.Intn(num + 1)
		fmt.Println(n)
	}
	return n
}
