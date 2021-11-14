package main

import (
	"math"

	"golang.org/x/net/context"
)

func (s *Server) solve1(ctx context.Context, max int) int {
	sum := 0
	for i := 3; i < max; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}

func (s *Server) solve2(ctx context.Context, max int) int {
	s1 := 1
	s2 := 2
	sum := 2

	for s2 < max {
		s1, s2 = s2, s1+s2
		if s2%2 == 0 && s2 < max {
			sum += s2
		}
	}

	return sum
}

func (s *Server) solve3(ctx context.Context, val int64) int64 {
	nums := primeFactors(val)
	return nums[len(nums)-1]
}

func (s *Server) solve4(ctx context.Context, val int64) int64 {
	start := math.Pow10(int(val - 1))
	end := math.Pow10(int(val))

	largest := int64(-1)
	for num1 := start; num1 < end; num1++ {
		for num2 := start; num2 < end; num2++ {
			res := int64(num1 * num2)
			if res > largest && isPalindrome(res) {
				largest = res
			}
		}
	}

	return largest
}
