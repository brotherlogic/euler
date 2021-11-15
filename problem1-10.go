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

func (s *Server) solve5(ctx context.Context, val int64) int64 {
	mval := int64(1)
	for i := int64(1); i <= val; i++ {
		mval *= i
	}
	for i := int64(1); i < mval; i++ {
		found := true
		for j := int64(2); j < val; j++ {
			if i%j != 0 {
				found = false
				break
			}
		}

		if found {
			return i
		}
	}

	return -1
}

func (s *Server) solve6(ctx context.Context, val int64) int64 {
	ssq := int64(0)
	sum := int64(0)
	for i := int64(1); i <= val; i++ {
		ssq += i * i
		sum += i
	}

	return sum*sum - ssq
}

func (s *Server) solve7(ctx context.Context, val int64) int64 {
	primes := getPrimes(21474836)
	return primes[val-1]
}
